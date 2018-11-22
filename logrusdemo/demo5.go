package main

import "github.com/sirupsen/logrus"

/*
 * Hook
 */

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.AddHook(&DefaultFieldsHook{})
	logrus.Info("test")
}

type DefaultFieldsHook struct {
}

func (df *DefaultFieldsHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}

func (df *DefaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
