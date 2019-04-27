package main

import (
	"github.com/sirikon/gonference/src/database"
	"github.com/sirikon/gonference/src/http"
	"github.com/sirikon/gonference/src/ioc"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := database.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := database.GetConnectionForDatabase("gonference")
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
