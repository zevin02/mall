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

// 更新头像
func UploadAvatar(c *gin.Context) {
	//请求的content-type
	file, fileHeader, _ := c.Request.FormFile("file") //从http请求中获取上传的文件，获取multipart的文件信息，file也是通过这个进行上传

	//fileheader是一个包含文件头信息的结构体
	fileSize := fileHeader.Size

	var userAvatar service.UserService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //根据token获得当前的各种信息
	if err := c.ShouldBind(&userAvatar); err == nil {
		//绑定成功
		//处理注册逻辑
		res := userAvatar.Post(c.Request.Context(), claims.ID, file, fileSize) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}
}

func SendEmail(c *gin.Context) {

	var sendEmail service.SendEmailService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //根据token获得当前的各种信息
	if err := c.ShouldBind(&sendEmail); err == nil {
		//绑定成功
		//处理注册逻辑
		res := sendEmail.Send(c.Request.Context(), claims.ID) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}
}
