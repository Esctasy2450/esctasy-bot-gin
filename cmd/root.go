package main

import (
	"esctasy-bot-gin/bot"
	"esctasy-bot-gin/controller"
	"github.com/gin-gonic/gin"
)

// Run 程序启动方法
func Run() {
	initLog()

	initBot()

	r := gin.Default()
	controller.SetupRouter(r)
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

// initBot 初始化机器人方法
func initBot() {
	bot.Init()
}
