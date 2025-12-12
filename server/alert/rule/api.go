package rule

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(g *gin.RouterGroup) {
	g.Group("/rule").
		POST("/page", PageAlertRule).
		POST("/add", AddAlertRule).
		POST("/modify", ModifyAlertRule).
		POST("/delete", DeleteAlertRule).
		POST("/submit", SubmitAlertRule).
		POST("/copy", CopyAlertRule)
}
