package response

import "github.com/gin-gonic/gin"

// Common c
type Common struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}
// Success s
func Success(data interface{}) *Common {
	return &Common{
		Data: data,
	}
}

// Failure f
func Failure(code int, msg string) *Common {
	return &Common{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// UnAuthority u
func UnAuthority() *Common {
	return &Common{
		ErrMsg:  "未授权的请求",
		ErrCode: 110,
	}
}

// RspAction r
func RspAction(c *gin.Context, d *Common) {
	c.JSON(200, d)
}

// ErrAction e
func ErrAction(c *gin.Context, msg string) {
	RspAction(c, Failure(-1, msg))
}

// SuccessAction s
func SuccessAction(c *gin.Context, data interface{}) {
	RspAction(c, Success(data))
}
// ErrRequestAction e
func ErrRequestAction(c *gin.Context) {
	RspAction(c, Failure(-1, "请求参数有误"))
}
