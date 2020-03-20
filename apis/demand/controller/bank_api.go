package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenkitter/tools/apis/demand/request"
	"github.com/stevenkitter/tools/models/tools"
	"github.com/stevenkitter/tools/response"
)

// BankList 支持的银行列表
func (ct *Controller) BankList(c *gin.Context) {
	banks := make([]*tools.Bank, 0)
	ct.d.Find(&banks)
	response.SuccessAction(c, banks)
}

// CardInfo 银行卡信息
func (ct *Controller) CardInfo(c *gin.Context) {
	var req request.CardInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ErrRequestAction(c)
	}
	result, err := ct.CardInfoBusiness(req.CardNo)
	if err != nil {
		response.ErrAction(c, err.Error())
		return
	}
	response.SuccessAction(c, result)
}

// CardNoValid 卡是否有效
func (ct *Controller) CardNoValid(c *gin.Context) {
	var req request.CardInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ErrRequestAction(c)
	}
	valid := ct.CardNoValidBusiness(req.CardNo)
	response.SuccessAction(c, valid)
}
