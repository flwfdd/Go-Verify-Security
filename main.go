package main

import (
	"awesomeProject/router"
	"awesomeProject/tool"
	"log"
)

func main() {
	// 解析app.json
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	// 配置数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Printf(err.Error())
	}

	// 配置路由
	app := router.SetRoute()

	// GO!
	err = app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		log.Printf(err.Error())
	}
}
