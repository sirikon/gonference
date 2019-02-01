package main

import (
	"log"

	"github.com/sirikon/gonference/postgres"
)

func main() {
	err := postgres.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
