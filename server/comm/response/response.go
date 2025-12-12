package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	MSG_SUCCESS_EMPTY   = ""
	MSG_SUCCESS         = "请求成功"
	MSG_SUCCESS_QUERY   = "查询成功"
	MSG_SUCCESS_MODIFY  = "修改成功"
	MSG_SUCCESS_DELEETE = "删除成功"
	MSG_SUCCESS_ADD     = "添加成功"
	MSG_SUCCESS_COPY    = "复制成功"
	MSG_SUCCESS_SUBMIT  = "提交成功"
)

const (
	CodeSuccess       = 2000
	CodeInvalidParams = 4001 // 参数错误
	CodeUnauthorized  = 4002 // 未认证
	CodeInternalError = 5000 // 内部错误
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

func SuccessQuery(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  MSG_SUCCESS_QUERY,
		Data: data,
	})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  MSG_SUCCESS_EMPTY,
		Data: data,
	})
}

func SuccessWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: nil,
	})
}

// 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ErrorWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeInternalError,
		Msg:  msg,
		Data: nil,
	})
}

func ErrorWithCode(c *gin.Context, code int, msg string) {
	Error(c, code, msg)
}

func InvalidParams(c *gin.Context, msg string) {
	Error(c, CodeInvalidParams, msg)
}

func Unauthorized(c *gin.Context) {
	Error(c, CodeUnauthorized, "未认证")
}
