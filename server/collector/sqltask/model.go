package sqltask

import "github.com/qwganker/boring/comm/request"

type SqlTaskPageReq struct {
	request.PageRequest
}

type SqlTaskAddReq struct {
	Remark             string `json:"remark" binding:"required"`
	DBType             string `json:"DBType" binding:"required"`
	DSN                string `json:"dsn" binding:"required"`
	SQL                string `json:"sql" binding:"required"`
	MetricDefine       string `json:"MetricDefine" binding:"required"`
	Cron               string `json:"cron" binding:"required"`
	PrometheusConfigID int64  `json:"PrometheusConfigID" binding:"required"`
	Enabled            string `json:"enabled"`
}

type SqlTaskDeleteReq struct {
	ID int64 `json:"id"`
}

type SqlTaskModifyReq struct {
	ID                 int64  `json:"id" binding:"required"`
	DBType             string `json:"DBType" binding:"required"`
	DSN                string `json:"dsn" binding:"required"`
	SQL                string `json:"sql" binding:"required"`
	MetricDefine       string `json:"MetricDefine" binding:"required"`
	Remark             string `json:"remark" binding:"required"`
	Cron               string `json:"cron" binding:"required"`
	PrometheusConfigID int64  `json:"PrometheusConfigID" binding:"required"`
	Enabled            string `json:"enabled"`
}

type SqlTaskCopyReq struct {
	ID int64 `json:"id"`
}

type SqlTaskRunReq struct {
	ID int64 `json:"id"`
}

type SqlTaskStopReq struct {
	ID int64 `json:"id"`
}

type SqlTaskRunTestReq struct {
	ID int64 `json:"id"`
}
