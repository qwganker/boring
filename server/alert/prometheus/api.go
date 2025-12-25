package prometheus

import "github.com/gin-gonic/gin"

func InitAPI(g *gin.RouterGroup) {

	svc := PrometheusService{}

	g.Group("/prometheus").
		POST("/listall", svc.ListAllPrometheusConfig).
		POST("/page", svc.PagePrometheusConfig).
		POST("/add", svc.AddPrometheusConfig).
		POST("/modify", svc.ModifyPrometheusConfig).
		POST("/delete", svc.DeletePrometheusConfig).
		POST("/submit", svc.SumbitPrometheusConfig).
		POST("/status", svc.CheckPrometheusStatus).
		POST("/copy", svc.CopyPrometheusConfig)
}
