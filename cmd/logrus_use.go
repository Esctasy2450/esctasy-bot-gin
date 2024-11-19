package main

import (
	"bytes"
	"encoding/json"
	"esctasy-bot-gin/constant"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var (
	logPath = "./logs"
	logFile = "gin.log"
)

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
	logrus.SetFormatter(&TimePriorityJSONFormatter{})
	// 设置日志记录级别
	logrus.SetLevel(logrus.DebugLevel)
}

type TimePriorityJSONFormatter struct{}

func (f *TimePriorityJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logData := LogData{}
	logData.Time = entry.Time.Format(constant.YYYYMMDDHHMMSS)
	logData.Level = entry.Level.String()
	logData.Message = entry.Message
	logData.Data = entry.Data

	// 序列化 JSON
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(logData); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

type LogData struct {
	Time    string        `json:"time"`
	Level   string        `json:"level"`
	Message string        `json:"message"`
	Data    logrus.Fields `json:"otherFields"`
}
