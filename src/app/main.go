package main

import (
	"gonference/pkg/infrastructure/config"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/database/migrator"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/infrastructure/web"
	"gonference/pkg/ioc"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Instance
	log.
		WithField("connectionString", cfg.Database.URL).
		WithField("port", cfg.Web.Port).
		Info("Starting")

	conn := client.GetDBClient(cfg.Database.URL)
	migrator.Migrate(conn)

	jobContext := ioc.CreateJobContext(conn)

	httpServer := web.Server{
		JobContext: jobContext,
	}

	err := httpServer.Run(cfg.Web.Port)
	if err != nil {
		log.Fatal(err)
	}
}
