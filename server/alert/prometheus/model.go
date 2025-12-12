package prometheus

import "github.com/qwganker/boring/comm/request"

type PrometheusConfigPageReq struct {
	request.PageRequest
	Remark string `json:"remark"`
}

type PrometheusConfigAddReq struct {
	Remark      string `json:"remark" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CtrlAddress string `json:"CtrlAddress" binding:"required"`
	Config      string `json:"config"`
	Rule        string `json:"rule"`
	Enabled     string `json:"enabled"`
}

type PrometheusConfigDeleteReq struct {
	ID int64 `json:"id"`
}

type PrometheusConfigModifyReq struct {
	ID          int64  `json:"id" binding:"required"`
	Remark      string `json:"remark"`
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CtrlAddress string `json:"CtrlAddress" binding:"required"`
	Config      string `json:"config"`
	Rule        string `json:"rule"`
	Enabled     string `json:"enabled"`
}

type PrometheusConfigCopyReq struct {
	ID int64 `json:"id"`
}

type PrometheusConfigSumbitReq struct {
	ID int64 `json:"id"`
}

type CheckPrometheusStatusReq struct {
	ID int64 `json:"id"`
}
