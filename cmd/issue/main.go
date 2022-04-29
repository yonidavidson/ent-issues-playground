package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect/sql/schema"

	_ "github.com/mattn/go-sqlite3"

	"issue/ent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(ctx, schema.WithAtlas(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
