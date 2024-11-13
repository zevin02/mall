package v1

import (
	"github.com/gin-gonic/gin"
	"mall/service"
	"net/http"
)

// gin框架的用户注册处理函数
// context提供了处理http请求和响应的上下文
func UserRegister(c *gin.Context) {
	var userRegister service.UserService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&userRegister); err == nil {
		//绑定成功
		//处理注册逻辑
		res := userRegister.Register(c.Request.Context()) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&userLogin); err == nil {
		//绑定成功
		//处理注册逻辑
		res := userLogin.Login(c.Request.Context()) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}
