package main

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"orm/demo/ent/model"
	"orm/demo/ent/model/migrate"
)

func main() {
	client, err := model.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres dbname=ent_demo password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
