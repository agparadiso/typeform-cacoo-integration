package main

import (
	"log"
	"net/http"

	"github.com/agparadiso/tfcacoo/pkg/server"
)

func main() {
	s := server.New()
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
