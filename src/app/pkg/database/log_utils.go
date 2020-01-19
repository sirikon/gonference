package database

import "gonference/pkg/infrastructure/logger"

func logSelect(log logger.Logger, query string) {
	log.
		WithField("sql", query).
		Debug("Executing select query")
}

func logMutation(log logger.Logger, query string) {
	log.
		WithField("sql", query).
		Info("Executing mutation query")
}
