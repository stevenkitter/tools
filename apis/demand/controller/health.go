package controller

import "github.com/gin-gonic/gin"

// HealthGetHandler sniff
func (ct *Controller) HealthGetHandler(c *gin.Context) {
	c.String(200, "ok")
}
