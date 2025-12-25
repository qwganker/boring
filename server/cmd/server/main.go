package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/alert"
	"github.com/qwganker/boring/collector"
	"github.com/qwganker/boring/collector/sqltask"
	"github.com/qwganker/boring/comm"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/conf"
	"github.com/qwganker/boring/job"
	"github.com/qwganker/boring/storage"
)

func loadJobTask() {

	job.GetJobScheduler().CleanAll()

	sqltaskList, err := sqltask.GetTaskList()
	if err != nil {
		panic(err)
	}

	for _, sqltask := range sqltaskList {
		schedId, err := job.GetJobScheduler().RegisterTask(table.JobTaskTypeSQL, sqltask.Cron, &sqltask)
		if err != nil {
			panic(err)
		}

		sqltask.SchedID = schedId
		storage.GetDBInstance().Save(sqltask)
	}

	go job.GetJobScheduler().Run()
}

var GinMode = gin.DebugMode

func main() {
	if GinMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	configPath := flag.String("config", "./conf.yaml", "配置文件路径")
	flag.Parse()

	conf.Load(*configPath)

	storage.Init(conf.GetConfig().DB)

	r := gin.Default()

	v1 := r.Group("/api/v1")

	alert.InitAPI(v1)
	collector.InitAPI(v1)
	job.InitAPI(v1)

	loadJobTask()

	address := fmt.Sprintf("%s:%d", conf.GetConfig().Server.Host, conf.GetConfig().Server.Port)

	log.Printf("Boring-server %s listening on %s\n", comm.BoringVersion, address)
	if err := r.Run(address); err != nil {
		panic(err)
	}

}
