package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/util"
	"mall/service"
	"net/http"
)

// 初始化所有需要秒杀的商品,将mysql的信息存入redis中
func InitSkillGoods(c *gin.Context) {
	var service service.SkillGoodsService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&service); err == nil {
		//绑定成功
		//处理注册逻辑
		res := service.InitSkillGoods(c.Request.Context()) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}

func SkillGoods(c *gin.Context) {
	var service service.SkillGoodsService //请求的参数
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))

	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&service); err == nil {
		//绑定成功
		//处理注册逻辑
		res := service.SkillGoods(c.Request.Context(), claim.ID) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}
