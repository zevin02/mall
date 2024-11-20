package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/util"
	"mall/service"
)

func OrderPay(c *gin.Context) {
	service := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.PayDown(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		util.LogrusObj.Infoln(err)
	}
}
