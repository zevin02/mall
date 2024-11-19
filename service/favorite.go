package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
)

// 查看到某个商品的id
type FavoritesService struct {
	ProductID  uint `form:"product_id" json:"product_id"`
	BossID     uint ` form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	model.BasePage
}

// 给当前用户的收藏夹里面新增一个商品
func (service *FavoritesService) Create(ctx context.Context, id uint) serializer.Response {
	var favorite model.Favorite //一个favortite
	var err error
	var user *model.User
	var boss *model.User
	var product *model.Product
	code := e.SUCCESS
	favoriteDao := dao.NewFavoriteDao(ctx)
	//使用用户id和商品id作为联合索引，查询看看这个收藏是否存在
	favorite, err = favoriteDao.GetFavoriteByUserIdProductId(id, service.ProductID)
	//根据id获得对应的用户信息
	userDao := dao.NewUserDao(ctx)
	user, _ = userDao.GetUserById(id)             //获取当前登陆用户的信息
	boss, _ = userDao.GetUserById(service.BossID) //获取boss的相关信息
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductID)
	if favorite == (model.Favorite{}) {
		//favorite为空，说明没有收藏过
		favorite = model.Favorite{
			UserId:    id,
			User:      *user,
			ProductId: service.ProductID,
			Product:   *product,
			BossId:    service.BossID,
			Boss:      *boss,
		}
		//把当前的收藏进行添加
		err = favoriteDao.CreateFavorite(&favorite)
		if err != nil {
			util.LogrusObj.Info(err)
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else {
		//already exist
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
