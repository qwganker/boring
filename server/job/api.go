package job

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(g *gin.RouterGroup) {

	svc := JobService{}

	g = g.Group("/job")
	g.Group("/jobtask").
		POST("/page", svc.Page)
	// POST("/add", svc.Add).
	// POST("/modify", svc.Modify).
	// POST("/delete", svc.Delete).
	// POST("/copy", svc.Copy).
	// POST("/run_test", svc.RunTest)
}
