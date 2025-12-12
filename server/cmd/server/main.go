package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/alert"
	"github.com/qwganker/boring/conf"
	"github.com/qwganker/boring/storage"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	configPath := flag.String("config", "./conf.yaml", "配置文件路径")
	flag.Parse()

	conf.Load(*configPath)

	if gin.Mode() != gin.ReleaseMode {
		conf.PrintLoadedConfig()
	}

	storage.Init(conf.GetConfig().DB)

	r := gin.Default()

	v1 := r.Group("/api/v1")

	alert.InitAPI(v1)

	address := fmt.Sprintf("%s:%d", conf.GetConfig().Server.Host, conf.GetConfig().Server.Port)
	log.Println("Listening and serving HTTP on " + address)

	if err := r.Run(address); err != nil {
		panic(err)
	}

}
