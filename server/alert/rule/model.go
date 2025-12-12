package rule

import "github.com/qwganker/boring/comm/request"

type AlertRulePageReq struct {
	request.PageRequest
	Type string `json:"type"`
}

type AlertRuleAddReq struct {
	Title              string `json:"title" binding:"required"`
	Level              string `json:"level" binding:"required"`
	Type               string `json:"type" binding:"required"`
	Source             string `json:"source"`
	PromQLRule         string `json:"PromQLRule" binding:"required"`
	Content            string `json:"content" binding:"required"`
	For                int64  `json:"for"`
	PromQLQuery        string `json:"PromQLQuery"`
	CustomLabels       string `json:"custom_labels"`
	NotifyID           int64  `json:"notifyID"`
	PrometheusConfigID int64  `json:"prometheusConfigId" binding:"required"`
	Enabled            string `json:"enabled"`
}

type AlertRuleDeleteReq struct {
	ID int64 `json:"id"`
}

type AlertRuleModifyReq struct {
	ID                 int64  `json:"id" binding:"required"`
	Title              string `json:"title"`
	Level              string `json:"level"`
	Type               string `json:"type"`
	Source             string `json:"source"`
	PromQLRule         string `json:"PromQLRule"`
	Content            string `json:"content"`
	For                int64  `json:"for"`
	PromQLQuery        string `json:"PromQLQuery"`
	CustomLabels       string `json:"CustomLabels"`
	NotifyID           int64  `json:"notifyID"`
	PrometheusConfigID int64  `json:"PrometheusConfigID" binding:"required"`
	Enabled            string `json:"enabled"`
}

type AlertRuleCopyReq struct {
	ID int64 `json:"id"`
}

type AlertRuleSumbitReq struct {
	ID int64 `json:"id"`
}
