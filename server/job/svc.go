package job

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/comm/request"
	"github.com/qwganker/boring/comm/response"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/storage"
)

type JobService struct {
}

func (j *JobService) Page(c *gin.Context) {
	var req JobTaskPageReq
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

	query := gormDB.WithContext(ctx).Model(&table.TJobTask{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		response.ErrorWithMsg(c, fmt.Sprintf("查询 JobTask 总数失败: %v", err))
		return
	}

	var items []table.TJobTask
	if total > 0 {
		if err := query.Order("id DESC").
			Offset(req.Offset()).
			Limit(req.Limit()).
			Find(&items).Error; err != nil {
			response.ErrorWithMsg(c, fmt.Sprintf("查询 JobTask 失败: %v", err))
			return
		}
	}

	response.SuccessWithData(c, request.NewPageResult(req.PageRequest, total, items))
}
