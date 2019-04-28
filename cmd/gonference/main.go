package main

import (
	"github.com/sirikon/gonference/src/database"
	"github.com/sirikon/gonference/src/http"
	"github.com/sirikon/gonference/src/ioc"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	conn, err := database.GetConnection(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(conn)
	if err != nil {
		log.Fatal(err)
	}

	serviceProvider := ioc.CreateServiceProvider(conn)

	httpServer := http.Server{
		ServiceProvider: serviceProvider,
	}

	err = httpServer.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
