package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gonference/pkg/utils"
	"os"
)

const generalLogfile = "logs/gonference.log"
const accessLogfile = "logs/gonference.access.log"

type Logger = logrus.FieldLogger
var Instance = createLogger(generalLogfile)
var InstanceForAccess = createLogger(accessLogfile)

func createLogger(logfile string) Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Hooks.Add(&ErrorHook{})

	file, err := os.OpenFile(logfile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755); utils.Check(err)
	logger.SetOutput(file)
	return logger
}

type ErrorHook struct {}

func (eh *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (eh *ErrorHook) Fire(entry *logrus.Entry) error {
	fmt.Println("ERROR: " + entry.Message)
	return nil
}
