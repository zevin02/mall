package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
	"strconv"
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
	var err error
	var user *model.User
	var boss *model.User
	var product *model.Product
	code := e.SUCCESS
	favoriteDao := dao.NewFavoriteDao(ctx)
	//使用用户id和商品id作为联合索引，查询看看这个收藏是否存在
	favorite, err := favoriteDao.GetFavoriteByUserIdProductId(id, service.ProductID)
	//根据id获得对应的用户信息
	userDao := dao.NewUserDao(ctx)
	user, _ = userDao.GetUserById(id)             //获取当前登陆用户的信息
	boss, _ = userDao.GetUserById(service.BossID) //获取boss的相关信息
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductID)
	if *favorite == (model.Favorite{}) {
		//favorite为空，说明没有收藏过
		favorite = &model.Favorite{
			UserId:    id,
			User:      *user,
			ProductId: service.ProductID,
			Product:   *product,
			BossId:    service.BossID,
			Boss:      *boss,
		}
		//把当前的收藏进行添加
		err = favoriteDao.CreateFavorite(favorite)
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

func (service *FavoritesService) Delete(ctx context.Context, uid uint, pid string) serializer.Response {
	var favorite *model.Favorite //一个favortite
	var err error
	code := e.SUCCESS
	//先检查一下当前商品是否存在
	favoriteDao := dao.NewFavoriteDao(ctx)
	//使用用户id和商品id作为联合索引，查询看看这个收藏是否存在
	pId, err := strconv.Atoi(pid)
	favorite, err = favoriteDao.GetFavoriteByUserIdProductId(uid, uint(pId))
	if err != nil {
		//
		util.LogrusObj.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//发现当前这个收藏已经存在，直接删除即可
	err = favoriteDao.DeleteFavorite(favorite)
	if err != nil {
		//
		util.LogrusObj.Info(err)
		code = e.ERROR
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

// 查看当前用户的所有收藏
func (service *FavoritesService) Show(ctx context.Context, id uint) serializer.Response {
	var favorites []*model.Favorite
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	favoriteDao := dao.NewFavoriteDao(ctx)
	//这个获得它的收藏夹
	if err := favoriteDao.DB.Model(&favorites).Preload("User").Where("user_id=?", id).Count(&total).Error; err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	favorites, err := favoriteDao.ListFavorite(id)

	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(total))

}
