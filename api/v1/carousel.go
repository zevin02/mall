package v1

import (
	"github.com/gin-gonic/gin"
	"mall/service"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&listCarousel); err == nil {
		//绑定成功
		//处理注册逻辑
		res := listCarousel.ListCarousel(c.Request.Context()) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}
