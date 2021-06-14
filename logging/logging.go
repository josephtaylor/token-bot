package logging

import (
	"github.com/sirupsen/logrus"
)

func ConfigureLogging(debug bool) {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
