package rule

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/alert/prometheus"
	"github.com/qwganker/boring/comm/constant"
	"github.com/qwganker/boring/comm/request"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/storage"
)

func PageAlertRule(c *gin.Context) {
	var req AlertRulePageReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}
	req.Normalize()

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	query := gormDB.WithContext(ctx).Model(&table.TAlertRule{})

	if req.Type != "" {
		query = query.Where(&table.TAlertRule{Type: req.Type})
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 t_alert_rule 总数失败: %v", err))
		return
	}

	var items []table.TAlertRule
	if total > 0 {
		if err := query.Order("id DESC").
			Offset(req.Offset()).
			Limit(req.Limit()).
			Find(&items).Error; err != nil {
			response.ErrorWithMsg(c, fmt.Sprintf("查询 t_alert_rule 失败: %v", err))
			return
		}

		// collect unique PrometheusConfigIDs
		idsMap := make(map[int64]struct{})
		var ids []int64
		for _, it := range items {
			if it.PrometheusConfigID != 0 {
				if _, ok := idsMap[it.PrometheusConfigID]; !ok {
					idsMap[it.PrometheusConfigID] = struct{}{}
					ids = append(ids, it.PrometheusConfigID)
				}
			}
		}

		// query all related Prometheus configs in batch
		cfgMap := make(map[int64]table.TPrometheusConfig)
		if len(ids) > 0 {
			var cfgs []table.TPrometheusConfig
			if err := gormDB.WithContext(ctx).Model(&table.TPrometheusConfig{}).
				Where("id IN ?", ids).
				Find(&cfgs).Error; err != nil {
				response.ErrorWithMsg(c, fmt.Sprintf("查询 t_prometheus_config 失败: %v", err))
				return
			}
			for _, c := range cfgs {
				cfgMap[c.ID] = c
			}
		}

		// build response items with attached prometheus config
		type AlertRuleWithConfig struct {
			table.TAlertRule
			PrometheusConfig *table.TPrometheusConfig `json:"PrometheusConfig,omitempty"`
		}

		var outItems []AlertRuleWithConfig
		outItems = make([]AlertRuleWithConfig, 0, len(items))
		for _, it := range items {
			ar := AlertRuleWithConfig{TAlertRule: it}
			if it.PrometheusConfigID != 0 {
				if cfg, ok := cfgMap[it.PrometheusConfigID]; ok {
					copyCfg := cfg
					ar.PrometheusConfig = &copyCfg
				}
			}
			outItems = append(outItems, ar)
		}

		response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, outItems))
		return
	}

	response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, items))
}

func AddAlertRule(c *gin.Context) {
	var req AlertRuleAddReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	rule := table.TAlertRule{
		Title:              req.Title,
		Level:              req.Level,
		Type:               req.Type,
		Source:             req.Source,
		PromQLRule:         req.PromQLRule,
		Content:            req.Content,
		For:                req.For,
		PromQLQuery:        req.PromQLQuery,
		CustomLabels:       req.CustomLabels,
		PrometheusConfigID: req.PrometheusConfigID,
		NotifyID:           req.NotifyID,
		Enabled:            req.Enabled,
	}

	if err := gormDB.WithContext(ctx).Create(&rule).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("添加告警规则失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_ADD)
}

func DeleteAlertRule(c *gin.Context) {
	var req AlertRuleDeleteReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var rule table.TAlertRule
	if err := gormDB.WithContext(ctx).First(&rule, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("告警规则不存在: %v", err))
		return
	}

	if err := gormDB.WithContext(ctx).Delete(&table.TAlertRule{}, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("删除告警规则失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_DELEETE)
}

func ModifyAlertRule(c *gin.Context) {
	var req AlertRuleModifyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var rule table.TAlertRule
	if err := gormDB.WithContext(ctx).First(&rule, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("告警规则不存在: %v", err))
		return
	}

	updates := make(map[string]interface{})
	updates["title"] = req.Title
	updates["level"] = req.Level
	updates["type"] = req.Type
	updates["source"] = req.Source
	updates["promql_rule"] = req.PromQLRule
	updates["content"] = req.Content
	updates["for"] = req.For
	updates["promql_query"] = req.PromQLQuery
	updates["custom_labels"] = req.CustomLabels
	updates["notify_id"] = req.NotifyID
	updates["enabled"] = req.Enabled
	updates["prometheus_config_id"] = req.PrometheusConfigID

	if len(updates) == 0 {
		response.SuccessWithMsg(c, response.MSG_SUCCESS)
		return
	}

	if err := gormDB.WithContext(ctx).Model(&table.TAlertRule{}).
		Where("id = ?", req.ID).
		Updates(updates).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("更新告警规则失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_MODIFY)
}

func CopyAlertRule(c *gin.Context) {
	var req AlertRuleCopyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var rule table.TAlertRule
	if err := gormDB.WithContext(ctx).First(&rule, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("告警规则不存在: %v", err))
		return
	}

	newRule := table.TAlertRule{
		Title:              rule.Title + "_COPY",
		Level:              rule.Level,
		Type:               rule.Type,
		Source:             rule.Source,
		PromQLRule:         rule.PromQLRule,
		Content:            rule.Content,
		For:                rule.For,
		PromQLQuery:        rule.PromQLQuery,
		CustomLabels:       rule.CustomLabels,
		NotifyID:           rule.NotifyID,
		PrometheusConfigID: rule.PrometheusConfigID,
		Enabled:            constant.Disabled,
	}

	if err := gormDB.WithContext(ctx).Create(&newRule).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("复制告警规则失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_COPY)
}

func SubmitAlertRule(c *gin.Context) {
	var req AlertRuleSumbitReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var rule table.TAlertRule
	if err := gormDB.WithContext(ctx).First(&rule, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("告警规则不存在: %v", err))
		return
	}

	ruleYaml, err := prometheus.RenderAlertRule(ctx, rule.PrometheusConfigID)
	if err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("渲染 Prometheus 规则失败: %v", err))
		return
	}

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, rule.PrometheusConfigID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 配置不存在: %v", err))
		return
	}

	config.Rule = ruleYaml

	if err := gormDB.WithContext(ctx).Save(&config).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("保存 Prometheus 规则失败: %v", err))
		return
	}

	ctrl := prometheus.NewPrometheusCtrl(c)

	if err := ctrl.RewritePrometheusRule(&config); err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Rewrite Prometheus Rule failed: %v", err))
		return
	}

	if err := ctrl.ReloadPrometheus(&config); err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Reload Prometheus failed: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_SUBMIT)
}
