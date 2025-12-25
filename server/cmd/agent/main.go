package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/comm"
)

var (
	webConfigFilePath        *string
	prometheusConfigFilePath *string
	ruleConfigFilePath       *string
)

var GinMode = gin.DebugMode

func main() {
	if GinMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	port := flag.Int("port", 7767, "specify the listening port")
	webConfigFilePath = flag.String("web.config.file", "./web-config.yml", "path to web config file")
	prometheusConfigFilePath = flag.String("config.file", "./prometheus.yml", "path to prometheus config file")
	ruleConfigFilePath = flag.String("rule.file", "./rule.yml", "path to prometheus rule config file")
	flag.Parse()

	webConfig, err := LoadWebConfig(*webConfigFilePath)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	r := gin.Default()
	g := r.Group("/")
	g.Use(BasicAuth(webConfig.BasicAuthUsers))

	g.POST("-/rewrite_rule", RewriteRuleFile)
	g.POST("-/rewrite_config", RewriteConfigFile)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	log.Printf("Boring-agent %s listening on %s\n", comm.BoringVersion, address)

	err = r.Run(address)
	if err != nil {
		panic(err)
	}
}
