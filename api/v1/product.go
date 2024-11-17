package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/util"
	"mall/service"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	//TODO
	form, _ := c.MultipartForm()                               //获取表单
	files := form.File["file"]                                 //当前商品可能上传了多张图片
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //根据token获得当前的各种信息
	createProdcutService := service.ProductService{}
	if err := c.ShouldBind(&createProdcutService); err == nil {
		//绑定成功
		//处理注册逻辑
		res := createProdcutService.Create(c.Request.Context(), claims.ID, files) //进行注册操作
		c.JSON(http.StatusOK, res)
	} else {
		//绑定失败
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln("create product error")
	}

}
