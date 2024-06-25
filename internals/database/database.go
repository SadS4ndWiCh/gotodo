package database

import (
	"context"
	"log"

	"github.com/SadS4ndWiCh/gotodo/ent"
	_ "github.com/mattn/go-sqlite3"
)

var (
	client *ent.Client
	err    error
)

func InitClient() {
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Panicf("Fatal error of sqlite\n %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func GetClient() *ent.Client {
	return client
}

func CloseClient() {
	client.Close()
}
