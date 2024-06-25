package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SadS4ndWiCh/gotodo/ent"
	"github.com/SadS4ndWiCh/gotodo/internals/database"
)

type Server struct {
	addr string
	db   *ent.Client
}

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	srv := &Server{
		addr: fmt.Sprintf(":%s", port),
		db:   database.GetClient(),
	}

	return &http.Server{
		Addr:    srv.addr,
		Handler: srv.Bootstrap(),
	}
}
