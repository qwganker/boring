package alerttype

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(g *gin.RouterGroup) {

	svc := NewAlertTypeService()

	g.Group("/alert_type").
		POST("/page", svc.Page).
		POST("/listall", svc.ListAll).
		POST("/add", svc.Add).
		POST("/modify", svc.Modify).
		POST("/delete", svc.Delete)
}
