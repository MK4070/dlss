package main

import (
	"log"

	"github.com/MK4070/dlss/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
