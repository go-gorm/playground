package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	//db connect with pgx
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		t.Errorf("Failed can't open db")
	}
	//create table events with pgx
	sql := `
	DROP TABLE IF EXISTS "events";
	DROP SEQUENCE IF EXISTS events_id_seq;
	CREATE SEQUENCE events_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
	CREATE TABLE "public"."events" (
		"id" integer DEFAULT nextval('events_id_seq') NOT NULL,
		"town" text NOT NULL,
		CONSTRAINT "events_pkey" PRIMARY KEY ("id")
	) WITH (oids = false);
	`
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		t.Errorf("Failed can't exec table events")
	}
	//insert event
	_, err = conn.Exec(context.Background(), "insert into events(town) values('Paris')")
	if err != nil {
		t.Errorf("Failed can't insert in table")
	}
	//close pgx db
	err = conn.Close(ctx)
	if err != nil {
		t.Errorf("Failed can't close pgx db")
	}
	//try to migrate
	err = DB.AutoMigrate(&Event{})
	if err != nil {
		t.Errorf("Failed can't migrate pgx to gorm")
	}
}
