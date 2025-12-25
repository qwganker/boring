package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/collector/sqltask"
	"github.com/qwganker/boring/comm"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/conf"
	"github.com/qwganker/boring/storage"
)

type JobTaskReq struct {
	ID      int64
	SchedID table.JobSchedID
	Type    table.JobTaskType
	Cron    string
	Payload string
	State   table.JobTaskState
}

func processSqltask(jobtaskReq *JobTaskReq) error {
	var t table.TSqlTask
	err := json.Unmarshal([]byte(jobtaskReq.Payload), &t)
	if err != nil {
		return err
	}

	log.Printf("Exec SQL task -> SqlTaskID %d jobtaskID %d, SchedID %s\n", t.ID, jobtaskReq.ID, jobtaskReq.SchedID)

	err = sqltask.ExecSqlTask(&t)
	if err != nil {
		return err
	}

	return nil
}

func process(c *gin.Context) {

	var jobtaskReq JobTaskReq
	if err := c.ShouldBindBodyWithJSON(&jobtaskReq); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	if jobtaskReq.Type == table.JobTaskTypeSQL {
		err := processSqltask(&jobtaskReq)
		if err != nil {
			log.Println(err.Error())
			response.ErrorWithMsg(c, err.Error())
			return
		}

		response.SuccessWithMsg(c, "")
		return
	}

	// TODO 接口类型任务
	if jobtaskReq.Type == table.JobTaskTypeHTTP {
	}

	response.SuccessWithMsg(c, "")
}

var GinMode = gin.DebugMode

func main() {

	if GinMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	configPath := flag.String("config", "./conf.yaml", "配置文件路径")
	flag.Parse()
	conf.Load(*configPath)
	// conf.PrintLoadedConfig()

	storage.Init(conf.GetConfig().DB)

	r := gin.Default()
	v1 := r.Group("/api/v1/job/worker")
	v1.POST("/process", process)

	address := fmt.Sprintf("%s:%d", conf.GetConfig().JobWorkerConfig.Host, conf.GetConfig().JobWorkerConfig.Port)
	log.Printf("Boring-jobworker %s listening on %s\n", comm.BoringVersion, address)
	if err := r.Run(address); err != nil {
		panic(err)
	}
}
