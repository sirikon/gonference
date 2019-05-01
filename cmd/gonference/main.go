package main

import (
	"github.com/sirikon/gonference/src/database"
	"github.com/sirikon/gonference/src/ioc"
	"github.com/sirikon/gonference/src/web"
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

	httpServer := web.Server{
		ServiceProvider: serviceProvider,
	}

	err = httpServer.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
