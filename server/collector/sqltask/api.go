package sqltask

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(g *gin.RouterGroup) {

	svc := NewSQLTaskService()

	g.Group("/sqltask").
		POST("/page", svc.Page).
		POST("/add", svc.Add).
		POST("/modify", svc.Modify).
		POST("/delete", svc.Delete).
		POST("/copy", svc.Copy).
		POST("/run_test", svc.RunTest)
}
