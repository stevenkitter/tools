package controller

import "github.com/stevenkitter/tools/middleware"

// Route 路由
func (ct *Controller) Route() {
	ct.g.Use(middleware.CORSMiddleware())

	ct.g.GET("/", ct.HealthGetHandler)

	auth := ct.g.Group("/auth")
	auth.Use(middleware.AuthorityMiddleware(ct.cache))
	auth.POST("/support/bankList", ct.BankList)

	auth.POST("/cardInfo", ct.CardInfo)
	auth.POST("/cardNo/valid", ct.CardNoValid)
}
