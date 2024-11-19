package v1

import (
	"github.com/gin-gonic/gin"
	"mall/service"
	"net/http"
)

func ListProductImg(c *gin.Context) {
	var listProductImg service.ListProductImg //请求的参数
	//绑定请求参数，尝试将请求中的数据json/表单数据绑定到userRegister中
	if err := c.ShouldBind(&listProductImg); err == nil {
		//绑定成功
		//处理注册逻辑
		res := listProductImg.List(c.Request.Context(), c.Param("id")) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
	}

}
