package main

import (
	"gonference/pkg/database"
	"gonference/pkg/infrastructure/config"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/ioc"
	"gonference/pkg/web"
)

func main() {

	cfg := config.ReadConfig()

	log := logger.Instance

	log.
		WithField("connectionString", cfg.Database.URL).
		WithField("port", cfg.Web.Port).
		Info("Starting")

	conn, err := database.GetConnection(cfg.Database.URL)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.GetConnectionPool(cfg.Database.URL)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(conn)
	if err != nil {
		log.Fatal(err)
	}

	jobContext := ioc.CreateJobContext(conn, pool)

	httpServer := web.Server{
		JobContext: jobContext,
	}

	err = httpServer.Run(cfg.Web.Port)
	if err != nil {
		log.Fatal(err)
	}
}
