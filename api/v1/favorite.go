package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/util"
	"mall/service"
)

func CreateFavorite(c *gin.Context) {
	service := service.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		util.LogrusObj.Infoln(err)
	}
}

// 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	service := service.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		util.LogrusObj.Infoln(err)
	}
}
