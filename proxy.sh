#!/bin/bash
set -e

# 颜色控制
RED='\033[0;31m'
NC='\033[0m' # No Color

# 检查 Docker 环境
if ! command -v docker &> /dev/null; then
    echo "未检测到 Docker 环境，正在安装..."
    curl -sSL https://get.docker.com/ | sh
    systemctl enable --now docker
    echo "Docker 安装完成。"
fi

# 默认值
DEFAULT_MTG_SECRET="ee4b18cadc815fceeda957cdddeeaa566e62696e672e636f6d"

# 交互式输入
read -p "请输入 MTG 密钥 (默认值: $DEFAULT_MTG_SECRET, 直接回车可使用): " MTG_CMD_VAR
MTG_CMD_VAR=${MTG_CMD_VAR:-$DEFAULT_MTG_SECRET}

while true; do
    read -p "请输入 Cloudflare API Token: " CF_TOKEN
    if [[ -z "$CF_TOKEN" ]]; then
        echo -e "${RED}Cloudflare API Token 不能为空，请重新输入。${NC}"
    elif [[ ${#CF_TOKEN} -lt 40 ]]; then
        echo -e "${RED}Cloudflare API Token 格式错误，请检查并重新输入。${NC}"
    else
        break
    fi
done

while true; do
    read -p "请输入域名 (多个请用逗号分隔, 例如: example.com,sub.example.com): " DOMAINS_INPUT
    if [[ -z "$DOMAINS_INPUT" ]]; then
        echo -e "${RED}域名不能为空，请重新输入。${NC}"
        continue
    fi

    # 检查域名格式
    IFS=',' read -ra CHECK_DOMAINS <<< "$DOMAINS_INPUT"
    ALL_VALID=true
    for domain in "${CHECK_DOMAINS[@]}"; do
        domain=$(echo "$domain" | xargs)
        if [[ ! $domain =~ ^([a-zA-Z0-9](([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)\.)+[a-zA-Z]{2,}$ ]]; then
            echo -e "${RED}域名格式不正确: $domain (示例: example.com)${NC}"
            ALL_VALID=false
            break
        fi
    done

    if [ "$ALL_VALID" = true ]; then
        break
    fi
done

read -p "请输入 V2Ray 端口 (默认 10625, 多个域名对应多组端口请用逗号分隔, 批量范围用-): " V2RAY_PORT_INPUT
V2RAY_PORT_INPUT=${V2RAY_PORT_INPUT:-"10625"}
IFS=',' read -ra V2RAY_PORTS_ARRAY <<< "$V2RAY_PORT_INPUT"

read -p "请输入 Singbox 起始端口 (默认 10808): " SINGBOX_PORT_START
SINGBOX_PORT_START=${SINGBOX_PORT_START:-10808}

# 解析域名并去除空格
IFS=',' read -ra DOMAINS <<< "$DOMAINS_INPUT"
DOMAINS=("${DOMAINS[@]// /}")

echo "正在生成配置文件..."

# --- 创建目录 ---
mkdir -p caddy/data caddy/config
echo "已创建 caddy 相关目录。"

# --- 生成 caddy.prod ---
cat > caddy.prod <<EOF
FROM caddy:builder-alpine AS builder
RUN xcaddy build --with github.com/caddy-dns/cloudflare

FROM caddy:alpine
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
EOF
echo "caddy.prod 已生成。"

# --- 生成 Caddyfile ---
cat > Caddyfile <<EOF
{
    email airplayx@gmail.com
    log {
        level ERROR
    }
}

:80 {
    root * /usr/share/caddy
    file_server
}
EOF

for domain in "${DOMAINS[@]}"; do
    cat >> Caddyfile <<EOF

$domain {
    encode gzip zstd
    log {
        level ERROR
    }
    tls {
        dns cloudflare {env.CLOUDFLARE_API_TOKEN}
    }

    @options {
        method OPTIONS
    }
    respond @options 204

    @blockUA {
        header_regexp User-Agent "(?i)baiduspider|360spider|Sogou web spider|Sosospider|YisouSpider|Bingbot|Googlebot|Scrapy|Curl|HttpClient|python-requests|Go-http-client|WinHttp|WebZIP|FetchURL|node-superagent|java/|FeedDemon|Jullo|JikeSpider|Indy Library|Alexa Toolbar|AskTbFXTV|AhrefsBot|CrawlDaddy|Java|Feedly|Apache-HttpAsyncClient|UniversalFeedParser|ApacheBench|Microsoft URL Control|Swiftbot|ZmEu|oBot|jaunty|Python-urllib|lightDeckReports Bot|YYSpider|DigExt|MJ12bot|heritrix|EasouSpider|Ezooms|BOT/0.1|YandexBot|FlightDeckReports|SemrushBot|Linguee Bot"
    }
    respond @blockUA 403
}
EOF
done
echo "Caddyfile 已生成。"

# --- 生成 singbox.json ---
export SINGBOX_PORT_START
export DOMAINS_JSON=$(printf '%s\n' "${DOMAINS[@]}" | python3 -c 'import sys, json; print(json.dumps([line.strip() for line in sys.stdin if line.strip()]))')

python3 -c '
import json, os

domains = json.loads(os.environ["DOMAINS_JSON"])
start_port = int(os.environ["SINGBOX_PORT_START"])

inbounds = []
for i, domain in enumerate(domains):
    inbounds.append({
        "type": "vless",
        "tag": f"vless-in-{i+1}",
        "listen": "::",
        "listen_port": start_port + i,
        "users": [
            {
                "uuid": "8f4b3a5b-47a5-45c6-957a-b3a5c678f9c1"
            }
        ],
        "tls": {
            "enabled": True,
            "server_name": domain,
            "key_path": f"/etc/sing-box/ssl/{domain}/{domain}.key",
            "certificate_path": f"/etc/sing-box/ssl/{domain}/{domain}.crt",
            "alpn": [
                "h2", 
                "http/1.1"
            ]
        },
        "transport": {
            "type": "http",
		    "path": "",
		    "host": [
                domain
            ]
        },
        "multiplex": {
            "enabled": True
        }
    })

config = {
    "log": {
        "level": "error"
    },
    "inbounds": inbounds,
    "outbounds": [
        {
            "type": "direct"
        }
    ]
}
print(json.dumps(config, indent=2))
' > singbox.json
echo "singbox.json 已生成。"

# --- 生成 v2ray.json ---
export V2RAY_PORTS_JSON=$(printf '%s\n' "${V2RAY_PORTS_ARRAY[@]}" | python3 -c 'import sys, json; print(json.dumps([line.strip() for line in sys.stdin if line.strip()]))')

python3 -c '
import json, os

domains = json.loads(os.environ["DOMAINS_JSON"])
ports = json.loads(os.environ["V2RAY_PORTS_JSON"])

inbounds = []
for i, domain in enumerate(domains):
    if i < len(ports):
        current_port = ports[i]
    else:
        last_port = ports[-1]
        if "-" in last_port:
             current_port = last_port
        else:
             current_port = str(int(last_port) + (i - len(ports) + 1))

    inbounds.append({
        "port": current_port, 
        "protocol": "vmess",
        "settings": {
            "clients": [
                {
                    "id": "11a93dc6-73a8-417b-a028-d42ce6fc2b95",
                    "alterId": 0
                }
            ]
        },
        "streamSettings": {
            "network": "h2",
            "security": "tls",
            "tlsSettings": {
                "certificates": [
                    {
                        "certificateFile": f"/etc/v2ray/ssl/{domain}/{domain}.crt",
                        "keyFile": f"/etc/v2ray/ssl/{domain}/{domain}.key"
                    }
                ]
            },
            "tcpSettings": {},
            "httpSettings": {
                "path": "/mVdcK4sX/"
            },
            "kcpSettings": {},
            "wsSettings": {},
            "quicSettings": {}
        },
        "domain": domain
    })

config = {
    "log": {
        "error": "/var/log/v2ray/error.log",
        "loglevel": "error"
    },
    "inbounds": inbounds,
    "outbounds": [
        {
            "protocol": "freedom",
            "settings": {
                "domainStrategy": "UseIP"
            }
        },
        {
            "protocol": "blackhole",
            "settings": {},
            "tag": "block"
        }
    ],
    "routing": {
        "rules": [
            {
                "type": "field",
                "ip": [
                    "0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "169.254.0.0/16",
                    "172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.168.0.0/16",
                    "198.18.0.0/15", "198.51.100.0/24", "203.0.113.0/24", "::1/128",
                    "fc00::/7", "fe80::/10"
                ],
                "outboundTag": "block"
            }
        ]
    }
}
print(json.dumps(config, indent=2))
' > v2ray.json
echo "v2ray.json 已生成。"

# --- 生成 docker-compose.yml ---
SB_VOLUMES=""
SB_PORTS=""
V2_VOLUMES=""

i=0
for domain in "${DOMAINS[@]}"; do
    domain=$(echo "$domain" | xargs)
    SB_VOLUMES+="      - ./caddy/data/caddy/certificates/acme-v02.api.letsencrypt.org-directory/$domain/:/etc/sing-box/ssl/$domain/\n"
    SB_VOLUMES+="      - ./caddy/data/caddy/certificates/acme.zerossl.com-v2-dv90/$domain/:/etc/sing-box/ssl/${domain}_zs/\n"
    
    current_sb_port=$((SINGBOX_PORT_START + i))
    SB_PORTS+="      - \"$current_sb_port:$current_sb_port\"\n"
    
    V2_VOLUMES+="      - ./caddy/data/caddy/certificates/acme-v02.api.letsencrypt.org-directory/$domain/:/etc/v2ray/ssl/$domain/\n"
    V2_VOLUMES+="      - ./caddy/data/caddy/certificates/acme.zerossl.com-v2-dv90/$domain/:/etc/v2ray/ssl/${domain}_zs/\n"
    i=$((i+1))
done

cat > docker-compose.yml <<EOF
services:
  mtg:
    image: p3terx/mtg
    container_name: mtg
    privileged: true
    restart: unless-stopped
    command: ["run", "$MTG_CMD_VAR"]
    network_mode: host

  caddy:
    build:
      context: .
      dockerfile: caddy.prod
    container_name: caddy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    environment:
      ACME_AGREE: true
      TZ: Asia/Shanghai
      CLOUDFLARE_API_TOKEN: $CF_TOKEN
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
      - ./caddy/data:/data
      - ./caddy/config:/config      
    networks:
      - app-network
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:80"]
      interval: 10s
      timeout: 5s
      retries: 5
      start_period: 10s

  singbox:
    image: ghcr.io/sagernet/sing-box
    container_name: singbox
    restart: unless-stopped
    depends_on:
      caddy:
        condition: service_healthy
    volumes:
$(printf "$SB_VOLUMES")
      - ./singbox.json:/etc/sing-box/config.json:ro
    command: -D /var/lib/sing-box -C /etc/sing-box/ run
    ports:
$(printf "$SB_PORTS")

  v2ray:
    image: jrohy/v2ray
    privileged: true
    container_name: v2ray
    restart: unless-stopped
    depends_on:
      caddy:
        condition: service_healthy
    volumes:
$(printf "$V2_VOLUMES")
      - ./v2ray.json:/etc/v2ray/config.json
    network_mode: host
      
networks:
  app-network:
    driver: bridge
EOF
echo "docker-compose.yml 已生成。"

# --- 智能启动与回滚逻辑 ---
# 1. 启动 Caddy
echo "正在启动 Caddy 并申请证书..."
docker compose build caddy
docker compose up -d caddy

# 2. 等待证书锁
echo "正在等待证书文件生成 (最长等待 150 秒)..."
FIRST_DOMAIN=$(echo "${DOMAINS[0]}" | xargs)
COUNTER=0
MAX_RETRIES=30 

while [ $COUNTER -lt $MAX_RETRIES ]; do
    # 只要在 data 目录下能找到对应域名后缀的 .crt 文件即可（兼容 LE 和 ZeroSSL 目录名）
    if find ./caddy/data -name "*${FIRST_DOMAIN}.crt" | grep -q "${FIRST_DOMAIN}"; then
        echo -e "\n检测到证书已就绪！"
        break
    fi
    printf "."
    sleep 5
    COUNTER=$((COUNTER + 1))
done

if [ $COUNTER -eq $MAX_RETRIES ]; then
    echo -e "\n${RED}警告：证书申请超时。如果后续启动失败，可能是由于证书尚未下发。${NC}"
fi

# 3. 启动其余服务
echo "正在启动 Singbox 和 V2Ray..."
if ! docker compose up -d; then
    echo -e "${RED}错误: 启动失败，执行自动回滚...${NC}"
    docker compose down 2>/dev/null || true
    rm -rf caddy/ caddy.prod Caddyfile singbox.json v2ray.json docker-compose.yml
    exit 1
fi

echo "部署成功！所有服务已进入运行状态。"
