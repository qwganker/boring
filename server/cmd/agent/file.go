package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RewriteRuleFile(c *gin.Context) {
	rewriteFile(c, *ruleConfigFilePath)
}

func RewriteConfigFile(c *gin.Context) {
	rewriteFile(c, *prometheusConfigFilePath)
}

func rewriteFile(c *gin.Context, path string) {
	yaml, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
	}

	err = os.WriteFile(path, yaml, 0664)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
	} else {
		c.Writer.WriteHeader(http.StatusOK)
	}
}
