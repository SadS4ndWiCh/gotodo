package main

import (
	"fmt"

	"github.com/SadS4ndWiCh/gotodo/internals/database"
	"github.com/SadS4ndWiCh/gotodo/internals/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.InitClient()
	defer database.CloseClient()

	srv := server.NewServer()

	fmt.Println("🔖 Server is running")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
