package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/stevenkitter/tools/apis/demand/request"
	"github.com/stevenkitter/tools/pack"
	"github.com/stevenkitter/tools/response"
)

// AuthorityMiddleware 效验
// nonce appId timestamp sign
func AuthorityMiddleware(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param request.AuthParam
		if err := c.ShouldBindQuery(&param); err != nil {
			_ = c.AbortWithError(200, err)
			return
		}
		key := fmt.Sprintf("app:secret:%s", param.AppID)
		secret := client.Get(key)
		if secret.Val() == "" {
			c.AbortWithStatusJSON(200, response.UnAuthority())
			return
		}
		if pack.Auth(param.Nonce, param.Timestamp, param.AppID, secret.Val(), param.Sign) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(200, response.UnAuthority())
			return
		}
	}
}
