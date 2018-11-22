package main

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	logrus.Info("test info")
}

func init() {
	logrus.AddHook(&ContextHook{})
}

type ContextHook struct {
}

func (hook *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 10)
	//表示自身栈中跳过6个栈帧数  并且把栈中剩余信息写入pc中。
	//0表示Callers自身的调用栈，1表示Callers所在的调用栈
	runtime.Callers(6, pc)
	//
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()

	funcName := frame.Func.Name()
	funcName = funcName[strings.LastIndexByte(funcName, filepath.Separator)+1:]
	fileName := frame.File[strings.LastIndexByte(frame.File, filepath.Separator)+1:]

	entry.Data["file"] = fileName
	entry.Data["func"] = funcName
	entry.Data["line"] = frame.Line
	entry.Data["ck"] = "ck168"

	return nil
}
