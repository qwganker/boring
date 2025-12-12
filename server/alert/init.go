package alert

import (
	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/alert/alerttype"
	"github.com/qwganker/boring/alert/prometheus"
	"github.com/qwganker/boring/alert/rule"
)

func InitAPI(g *gin.RouterGroup) {
	group := g.Group("/alert")

	rule.InitAPI(group)
	prometheus.InitAPI(group)
	alerttype.InitAPI(group)
}
