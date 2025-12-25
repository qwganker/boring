package collector

import (
	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/collector/sqltask"
)

func InitAPI(g *gin.RouterGroup) {
	group := g.Group("/collector")

	sqltask.InitAPI(group)
}
