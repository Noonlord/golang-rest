package db

import (
	"context"
	"log"
	"os"

	"api-ent/ent"

	_ "github.com/lib/pq"
)

var entDb *ent.Client
var connStr string

func migrateDb(client *ent.Client) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func GetDb() *ent.Client {
	return entDb
}

func OpenDb() {
	envConnStr := os.Getenv("DATABASE_URL")
	if envConnStr != "" {
		connStr = envConnStr
	} else {
		connStr = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
	}

	client, err := ent.Open("postgres", connStr)
	migrateDb(client)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	entDb = client
}
