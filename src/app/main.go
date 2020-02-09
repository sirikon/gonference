package main

import (
	"gonference/pkg/database"
	"gonference/pkg/database/migrator"
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

	err = migrator.Migrate(conn)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.GetConnectionPool(cfg.Database.URL)
	if err != nil {
		log.Fatal(err)
	}

	jobContext := ioc.CreateJobContext(pool)

	httpServer := web.Server{
		JobContext: jobContext,
	}

	err = httpServer.Run(cfg.Web.Port)
	if err != nil {
		log.Fatal(err)
	}
}
