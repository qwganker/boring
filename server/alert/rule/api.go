package rule

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(g *gin.RouterGroup) {

	svc := RuleService{}
	g.Group("/rule").
		POST("/page", svc.PageAlertRule).
		POST("/add", svc.AddAlertRule).
		POST("/modify", svc.ModifyAlertRule).
		POST("/delete", svc.DeleteAlertRule).
		POST("/submit", svc.SubmitAlertRule).
		POST("/copy", svc.CopyAlertRule)
}
