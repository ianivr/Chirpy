package main

import (
	"log"

	"github.com/ianivr/chirpy/internal/server"
)

func main() {
	log.Fatal(server.Start())
}
