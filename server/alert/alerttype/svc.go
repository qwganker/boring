package alerttype

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/comm/request"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/storage"
)

type AlertTypeService struct {
}

func NewAlertTypeService() *AlertTypeService {
	return &AlertTypeService{}
}

func (*AlertTypeService) Page(c *gin.Context) {
	var req AlertTypePageReq
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

	query := gormDB.WithContext(ctx).Model(&table.TAlertType{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 t_alert_type 总数失败: %v", err))
		return
	}

	var items []table.TAlertType
	if total > 0 {
		if err := query.Order("id DESC").
			Offset(req.Offset()).
			Limit(req.Limit()).
			Find(&items).Error; err != nil {
			response.ErrorWithMsg(c, fmt.Sprintf("查询 t_alert_type 失败: %v", err))
			return
		}
	}

	response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, items))
}

func (*AlertTypeService) ListAll(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var items []table.TAlertType
	if err := gormDB.WithContext(ctx).Find(&items).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 t_alert_type 失败: %v", err))
		return
	}

	response.SuccessWithData(c, items)
}

func (*AlertTypeService) Add(c *gin.Context) {
	var req AlertTypeAddReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	config := table.TAlertType{
		Name: req.Name,
		Code: req.Code,
	}

	if err := gormDB.WithContext(ctx).Create(&config).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("添加 t_alert_type 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_ADD)
}

func (*AlertTypeService) Delete(c *gin.Context) {
	var req AlertTypeDeleteReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var config table.TAlertType
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("t_alert_type 配置不存在: %v", err))
		return
	}

	// 删除记录
	if err := gormDB.WithContext(ctx).Delete(&table.TAlertType{}, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("删除 t_alert_type 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_DELEETE)
}

func (*AlertTypeService) Modify(c *gin.Context) {
	var req AlertTypeModifyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var config table.TAlertType
	if err := gormDB.WithContext(ctx).First(&config, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("t_alert_type 配置不存在: %v", err))
		return
	}

	config.Name = req.Name
	config.Code = req.Code

	if err := gormDB.WithContext(ctx).Model(&table.TAlertType{}).
		Where("id = ?", req.ID).
		Updates(config).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("更新 t_alert_type 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_MODIFY)
}
