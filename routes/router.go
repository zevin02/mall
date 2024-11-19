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
		v1.POST("carousels", api.ListCarousel) //展示所有的商品信息

		v1.GET("products", api.ListProduct)    //展示所有的商品信息
		v1.GET("product/:id", api.ShowProduct) //展示所有的商品信息
		v1.GET("imgs/:id", api.ListProductImg) //展示所有的商品信息
		//
		authed := v1.Group("/") //需要登陆保护
		//意味着authed这个路由组下的所有请求都需要经过JWTAuth中间件
		authed.Use(middleware.JWTAuth()) //使用jwt中间件
		{
			authed.POST("user/update", api.UserUpdate)
			authed.POST("user/avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail) //发送邮件
			authed.POST("user/valid-email", api.ValidEmail)  //发送邮件

			//显示金额
			authed.POST("user/show_money", api.ShowMoney)    //展示金额
			authed.POST("product/create", api.CreateProduct) //创建商品
			authed.POST("product/search", api.SearchProduct) //创建商品

		}

	}
	return r
}
