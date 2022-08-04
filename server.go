package main

import (
	"github.com/gorilla/mux"
)

type server struct {
	db     *database
	router *mux.Router
}

func newServer() *server {
	server := &server{}
	server.routes()
	return server
}
