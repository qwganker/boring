package prometheus

import "github.com/gin-gonic/gin"

func InitAPI(g *gin.RouterGroup) {
	g.Group("/prometheus").
		POST("/listall", ListAllPrometheusConfig).
		POST("/page", PagePrometheusConfig).
		POST("/add", AddPrometheusConfig).
		POST("/modify", ModifyPrometheusConfig).
		POST("/delete", DeletePrometheusConfig).
		POST("/submit", SumbitPrometheusConfig).
		POST("/status", CheckPrometheusStatus).
		POST("/copy", CopyPrometheusConfig)
}
