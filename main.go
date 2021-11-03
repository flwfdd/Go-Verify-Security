package main

import (
	"awesomeProject/controller"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost+":"+cfg.AppPort)
}

func registerRouter(router *gin.Engine)  {
	new(controller.RegisterController).Router(router)
}