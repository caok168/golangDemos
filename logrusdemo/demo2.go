package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{DisableColors: false, FullTimestamp: false})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
