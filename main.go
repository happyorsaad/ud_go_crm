package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := server{
		router: mux.NewRouter(),
		db:     NewDatabase(),
	}
	s.routes()

	seedDatabase(s.db)

	fmt.Println("Server is starting on port 12000...")
	return http.ListenAndServe(":12000", s.router)
}
