package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

func main() {
	testHook()
}

func testHook() {
	hook := graylog.NewGraylogHook(
		"localhost",
		map[string]interface{}{"service": "IVS"},
	)
	logrus.AddHook(hook)

	logrus.Info("test")

}
