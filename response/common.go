package response

import "github.com/gin-gonic/gin"

type Common struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Common {
	return &Common{
		Data: data,
	}
}

func Failure(code int, msg string) *Common {
	return &Common{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func UnAuthority() *Common {
	return &Common{
		ErrMsg:  "未授权的请求",
		ErrCode: 110,
	}
}

func RspAction(c *gin.Context, d *Common) {
	c.JSON(200, d)
}

func ErrAction(c *gin.Context, msg string) {
	RspAction(c, Failure(-1, msg))
}

func SuccessAction(c *gin.Context, data interface{}) {
	RspAction(c, Success(data))
}

func ErrRequestAction(c *gin.Context) {
	RspAction(c, Failure(-1, "请求参数有误"))
}
