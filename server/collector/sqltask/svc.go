package sqltask

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/alert/prometheus"
	"github.com/qwganker/boring/comm/constant"
	"github.com/qwganker/boring/comm/request"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/job"
	"github.com/qwganker/boring/storage"
	"github.com/qwganker/boring/utils/promclient"
	"github.com/qwganker/boring/utils/series"
)

type SQLTaskService struct {
}

func NewSQLTaskService() *SQLTaskService {
	return &SQLTaskService{}
}

func (s *SQLTaskService) Page(c *gin.Context) {
	var req SqlTaskPageReq
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

	query := gormDB.WithContext(ctx).Model(&table.TSqlTask{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 SqlTask 总数失败: %v", err))
		return
	}

	var items []table.TSqlTask
	if total > 0 {
		if err := query.Order("id DESC").
			Offset(req.Offset()).
			Limit(req.Limit()).
			Find(&items).Error; err != nil {
			response.ErrorWithMsg(c, fmt.Sprintf("查询 SqlTask 失败: %v", err))
			return
		}
	}

	response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, items))
}

func (s *SQLTaskService) Add(c *gin.Context) {
	var req SqlTaskAddReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if series.ValidateMetricDefine(req.MetricDefine) != nil {
		response.ErrorWithMsg(c, "指标配置格式不正确")
		return
	}

	gormDB := storage.GetDBInstance()

	task := table.TSqlTask{
		SQL:                req.SQL,
		MetricDefine:       req.MetricDefine,
		DBType:             req.DBType,
		DSN:                req.DSN,
		Remark:             req.Remark,
		Cron:               req.Cron,
		SchedID:            "",
		PrometheusConfigID: req.PrometheusConfigID,
		Enabled:            req.Enabled,
	}

	if err := gormDB.WithContext(ctx).Create(&task).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("添加 SqlTask 配置失败: %v", err))
		return
	}

	if task.Enabled == constant.Enabled {
		var err error
		task.SchedID, err = job.GetJobScheduler().RegisterTask(table.JobTaskTypeSQL, task.Cron, &task)
		if err != nil {
			response.ErrorWithMsg(c, err.Error())
			return
		}

		gormDB.Save(task)
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_ADD)
}

func (s *SQLTaskService) Modify(c *gin.Context) {
	var req SqlTaskModifyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if series.ValidateMetricDefine(req.MetricDefine) != nil {
		response.ErrorWithMsg(c, "指标配置格式不正确")
		return
	}

	gormDB := storage.GetDBInstance()

	var task table.TSqlTask
	if err := gormDB.WithContext(ctx).First(&task, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("SqlTask 配置不存在: %v", err))
		return
	}

	task.DBType = req.DBType
	task.DSN = req.DSN
	task.SQL = req.SQL
	task.MetricDefine = req.MetricDefine
	task.Remark = req.Remark
	task.Cron = req.Cron
	task.PrometheusConfigID = req.PrometheusConfigID
	task.Enabled = req.Enabled

	if err := gormDB.WithContext(ctx).Model(&table.TSqlTask{}).
		Where("id = ?", req.ID).
		Updates(task).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("更新 SqlTask 配置失败: %v", err))
		return
	}

	if task.Enabled == constant.Enabled {
		var err error
		task.SchedID, err = job.GetJobScheduler().ReRegisterTask(task.SchedID, table.JobTaskTypeSQL, task.Cron, &task)
		if err != nil {
			response.ErrorWithMsg(c, err.Error())
			return
		}
		gormDB.Save(task)

	} else {
		job.GetJobScheduler().UnRegisterTask(task.SchedID)
		task.SchedID = ""
		gormDB.Save(task)
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_MODIFY)
}

func (s *SQLTaskService) Delete(c *gin.Context) {
	var req SqlTaskDeleteReq
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
	var task table.TSqlTask
	if err := gormDB.WithContext(ctx).First(&task, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("SqlTask 配置不存在: %v", err))
		return
	}

	job.GetJobScheduler().UnRegisterTask(task.SchedID)

	// 删除记录
	if err := gormDB.WithContext(ctx).Delete(&table.TSqlTask{}, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("删除 SqlTask 配置失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_DELEETE)
}

func (s *SQLTaskService) Copy(c *gin.Context) {
	var req SqlTaskCopyReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var sqlTask table.TSqlTask
	if err := gormDB.WithContext(ctx).First(&sqlTask, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("SqlTask 不存在: %v", err))
		return
	}

	newSqlTask := table.TSqlTask{
		DBType:             sqlTask.DBType,
		DSN:                sqlTask.DSN,
		SQL:                sqlTask.SQL,
		MetricDefine:       sqlTask.MetricDefine,
		Remark:             sqlTask.Remark + "_copy",
		Cron:               sqlTask.Cron,
		SchedID:            "",
		PrometheusConfigID: sqlTask.PrometheusConfigID,
		Enabled:            constant.Disabled,
	}

	if err := gormDB.WithContext(ctx).Create(&newSqlTask).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("复制 SqlTask 失败: %v", err))
		return
	}

	response.SuccessWithMsg(c, response.MSG_SUCCESS_COPY)
}

func (s *SQLTaskService) RunTest(c *gin.Context) {
	var req SqlTaskRunTestReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gormDB := storage.GetDBInstance()

	var task table.TSqlTask
	if err := gormDB.WithContext(ctx).First(&task, req.ID).Error; err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}

	err := ExecSqlTask(&task)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, "测试成功, 指标已推送到 Prometheus")
}

func GetTaskList() ([]table.TSqlTask, error) {
	list := []table.TSqlTask{}
	if err := storage.GetDBInstance().WithContext(context.Background()).Where("enabled = ?", constant.Enabled).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func ExecSqlTask(task *table.TSqlTask) error {

	sqlExecutor := newSQLExecutor(task.DBType, task.DSN, 5)

	result, err := sqlExecutor.exec(task.SQL)
	if err != nil {
		return err
	}

	mdefine, err := series.ParseMetricDefine(task.MetricDefine)
	if err != nil {
		return err
	}

	prom, err := prometheus.QueryPrometheusConfigByID(context.Background(), task.PrometheusConfigID)
	if err != nil {
		return err
	}

	promSeries, err := series.NewJsonConverter(result, mdefine).EncodeTimeSeries()
	if err != nil {
		return err
	}

	pusher := promclient.NewPusher(prom.Address, prom.Username, prom.Password, 5)
	err = pusher.Push(promSeries)
	if err != nil {
		return err
	}

	return nil
}
