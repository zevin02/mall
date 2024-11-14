package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/util"
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
		c.JSON(http.StatusOK, res)                        //gin框架会自动将res转化成JSON格式，res中的字段会根据标签转化成相应的json 键值对
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

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //根据token获得当前的各种信息
	if err := c.ShouldBind(&userUpdate); err == nil {
		//绑定成功
		//处理注册逻辑
		res := userUpdate.Update(c.Request.Context(), claims.ID) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}
