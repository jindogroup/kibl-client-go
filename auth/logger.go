package auth

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Logger() (log *logrus.Entry) {
	host, err := os.Hostname()
	if err != nil {
		host = os.Getenv("HOST")
	}
	log = logrus.WithField("host", host)
	log.Logger.SetLevel(logrus.DebugLevel)
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	log.Logger.SetReportCaller(true)
	return
}
