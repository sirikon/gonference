package main

import (
	"github.com/sirikon/gonference/ioc"
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

	serviceProvider := ioc.CreateServiceProvider(conn)

	httpServer := http.Server{
		ServiceProvider: serviceProvider,
	}

	err = httpServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
