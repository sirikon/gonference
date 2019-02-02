package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/sirikon/gonference/http"
	"github.com/sirikon/gonference/postgres"
)

func main() {
	err := postgres.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := postgres.GetConnectionForDatabase("gonference")
	if err != nil {
		log.Fatal(err)
	}

	talkRepository := &postgres.TalkRepository{
		DB: conn,
	}

	httpServer := http.Server{
		TalkRepository: talkRepository,
	}

	err = httpServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
