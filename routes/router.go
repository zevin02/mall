package routes

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/middleware"
	"net/http"
)

// 路由配置,各种前端请求路由走到这里，进行具体的操作
func NewRouter() *gin.Engine {
	r := gin.Default() //创建一个gin实例，包括日志和回复
	//使用自定义的cors中间件，处理跨域请求
	r.Use(middleware.Cors())
	//设置静态文件服务器，将/static映射到./static
	r.StaticFS("/static", http.Dir("./static"))
	//创建一个api路由组，路径前缀api/v1
	v1 := r.Group("api/v1")

	{
		//定义一个get请求,路径为/ping
		v1.GET("/ping", func(c *gin.Context) {
			//返回json响应，状态码为200,内容success
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
	}
	return r
}
