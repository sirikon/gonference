package main

import (
	"gonference/pkg/database"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/ioc"
	"gonference/pkg/web"
	"os"
)

func main() {
	log := logger.Instance
	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	log.
		WithField("connectionString", connectionString).
		WithField("port", port).
		Info("Starting")

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
