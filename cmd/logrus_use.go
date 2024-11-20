package main

import (
	"bytes"
	"esctasy-bot-gin/constant"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
)

var (
	logPath = "./logs"
	logFile = "gin.log"
)

// initLog 初始化日志配置
func initLog() {
	//打开文件
	logFileName := path.Join(logPath, logFile)
	fileWriter, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	//设置日志输出到文件
	logrus.SetOutput(io.MultiWriter(os.Stdout, fileWriter))
	logrus.SetReportCaller(true)
	// 设置日志输出格式
	logrus.SetFormatter(&formatter{})
	// 设置日志记录级别
	logrus.SetLevel(logrus.DebugLevel)
}

type formatter struct{}

func (m *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format(constant.YYYYMMDDHHMMSS)
	var newLog string

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		funName := filepath.Base(entry.Caller.Function)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d:%s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, funName, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}
