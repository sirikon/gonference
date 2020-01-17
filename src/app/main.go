package main

import (
	log "github.com/sirupsen/logrus"
	"gonference/pkg/database"
	"gonference/pkg/ioc"
	"gonference/pkg/web"
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

	jobContext := ioc.CreateJobContext(conn)

	httpServer := web.Server{
		JobContext: jobContext,
	}

	err = httpServer.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
