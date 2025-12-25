package prometheus

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/comm/constant"
	"github.com/qwganker/boring/comm/request"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/storage"
)

type PrometheusService struct {
}

func QueryPrometheusConfigByID(ctx context.Context, id int64) (*table.TPrometheusConfig, error) {
	gormDB := storage.GetDBInstance()

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, id).Error; err != nil {
		return nil, fmt.Errorf("查询 Prometheus 配置失败: %v", err)
	}

	return &config, nil
}

func (p *PrometheusService) QueryPrometheusConfigByID(ctx context.Context, id int64) (*table.TPrometheusConfig, error) {
	gormDB := storage.GetDBInstance()

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, id).Error; err != nil {
		return nil, fmt.Errorf("查询 Prometheus 配置失败: %v", err)
	}

	return &config, nil
}

func (p *PrometheusService) ListAllPrometheusConfig(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var items []table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).Find(&items).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 t_prometheus_config 失败: %v", err))
		return
	}

	response.SuccessWithData(c, items)
}

func (p *PrometheusService) PagePrometheusConfig(c *gin.Context) {
	var req PrometheusConfigPageReq
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

	query := gormDB.WithContext(ctx).Model(&table.TPrometheusConfig{})

	if req.Remark != "" {
		query = query.Where("remark LIKE ?", "%"+req.Remark+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 t_prometheus_config 总数失败: %v", err))
		return
	}

	var cfgs []table.TPrometheusConfig
	if total > 0 {
		if err := query.Order("id DESC").
			Offset(req.Offset()).
			Limit(req.Limit()).
			Find(&cfgs).Error; err != nil {
			response.ErrorWithMsg(c, fmt.Sprintf("查询 t_prometheus_config 失败: %v", err))
			return
		}
	}

	type PrometheusConfigWithStatus struct {
		table.TPrometheusConfig
		Status string `json:"Status"`
	}

	var cfgWithStatus []PrometheusConfigWithStatus
	for _, cfg := range cfgs {
		if err := p.checkStatus(c, cfg); err == nil {
			cfgWithStatus = append(cfgWithStatus, PrometheusConfigWithStatus{
				TPrometheusConfig: cfg,
				Status:            "Normal",
			})
			continue
		} else {
			cfgWithStatus = append(cfgWithStatus, PrometheusConfigWithStatus{
				TPrometheusConfig: cfg,
				Status:            "Error",
			})
			continue
		}
	}

	response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, cfgWithStatus))
}

func (p *PrometheusService) AddPrometheusConfig(c *gin.Context) {
	var req PrometheusConfigAddReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	config := table.TPrometheusConfig{
		Remark:      req.Remark,
		Address:     req.Address,
		Username:    req.Username,
		Password:    req.Password,
		CtrlAddress: req.CtrlAddress,
		Config:      req.Config,
		Rule:        req.Rule,
		Enabled:     req.Enabled,
	}

	if err := gormDB.WithContext(ctx).Create(&config).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("添加 Prometheus 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_ADD)
}

func (p *PrometheusService) DeletePrometheusConfig(c *gin.Context) {
	var req PrometheusConfigDeleteReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	// 检查记录是否存在
	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 配置不存在: %v", err))
		return
	}

	// 删除记录
	if err := gormDB.WithContext(ctx).Delete(&table.TPrometheusConfig{}, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("删除 Prometheus 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_DELEETE)
}

func (p *PrometheusService) ModifyPrometheusConfig(c *gin.Context) {
	var req PrometheusConfigModifyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 配置不存在: %v", err))
		return
	}

	config.Remark = req.Remark
	config.Address = req.Address
	config.Username = req.Username
	config.Password = req.Password
	config.CtrlAddress = req.CtrlAddress
	config.Config = req.Config
	config.Rule = req.Rule
	config.Enabled = req.Enabled

	if err := gormDB.WithContext(ctx).Model(&table.TPrometheusConfig{}).
		Where("id = ?", req.ID).
		Updates(config).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("更新 Prometheus 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_MODIFY)
}

func (p *PrometheusService) CopyPrometheusConfig(c *gin.Context) {
	var req PrometheusConfigCopyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var cfg table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&cfg, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("告警规则不存在: %v", err))
		return
	}

	newPrometheusConfig := table.TPrometheusConfig{
		Remark:      cfg.Remark + "_COPY",
		Address:     cfg.Address,
		Username:    cfg.Username,
		Password:    cfg.Password,
		CtrlAddress: cfg.CtrlAddress,
		Config:      "",
		Rule:        "",
		Enabled:     constant.Disabled,
	}

	if err := gormDB.WithContext(ctx).Create(&newPrometheusConfig).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("复制普米配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_COPY)
}

func (p *PrometheusService) SumbitPrometheusConfig(c *gin.Context) {
	var req PrometheusConfigSumbitReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 配置不存在: %v", err))
		return
	}

	ctrl := NewPrometheusCtrl(c)

	if err := ctrl.RewritePrometheusConfig(&config); err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("提交 Prometheus 配置失败: %v", err))
		return
	}

	if err := ctrl.RewritePrometheusRule(&config); err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("提交 Prometheus 规则失败: %v", err))
		return
	}

	if err := ctrl.ReloadPrometheus(&config); err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Reload Prometheus failed: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_SUBMIT)
}

func (p *PrometheusService) CheckPrometheusStatus(c *gin.Context) {
	var req CheckPrometheusStatusReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var config table.TPrometheusConfig
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 配置不存在: %v", err))
		return
	}

	err := p.checkStatus(c, config)
	if err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("Prometheus 状态异常: %v", err))
		return
	}

	response.SuccessWithMsg(c, "")
}

func (p *PrometheusService) checkStatus(c *gin.Context, config table.TPrometheusConfig) error {
	ctrl := NewPrometheusCtrl(c)

	if err := ctrl.CheckPrometheusStatus(&config); err != nil {
		return err
	}

	return nil
}
