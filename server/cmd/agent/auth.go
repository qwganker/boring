package main

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuth(users map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			setError(c, "认证失败, Header未设置 Authorization")
			return
		}

		if !strings.HasPrefix(auth, "Basic ") {
			setError(c, "认证失败, Header未设置 Basic")
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			setError(c, "认证失败, Authorization 格式错误")
			return
		}
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 {
			setError(c, "认证失败, Header未设置")
			return
		}

		username := pair[0]
		password := pair[1]

		if hash, exists := users[username]; exists {
			// bcrypt.CompareHashAndPassword 只会在完全匹配时返回 nil
			if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err == nil {
				c.Next()
				return
			}
		}

		setError(c, "认证失败, 错误的账号/密码")
	}
}

func setError(c *gin.Context, text string) {
	c.Writer.WriteHeader(http.StatusUnauthorized)
	c.Writer.Write([]byte(text))
}
