package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.TraceLevel)
	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
	log.Panic("panic")
}
