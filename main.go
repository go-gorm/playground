package main

import (
	"fmt"
	"gorm.io/gorm/schema"
)

func main() {
	ns := schema.NamingStrategy{}
	// Results in http_s_certificates, should be https_certificates
	fmt.Println(ns.TableName("HTTPSCertificate"))
	// Results in fb_ipassports, should be fbi_passports
	fmt.Println(ns.TableName("FBIPassport"))
	// Works as intended, resulting in cia_passports
	fmt.Println(ns.TableName("CIAPassport"))
}
