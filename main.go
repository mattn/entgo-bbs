package main

import (
	"context"
	"log"

	"github.com/mattn/entgo-bbs/ent"

	_ "github.com/mattn/go-sqlite3"
)

type Payload struct {
	Content string `json:"content"`
}

func main() {
	client, err := ent.Open("sqlite3", "file:bbs.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
