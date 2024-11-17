package service

import (
	"context"
	"mall/dao"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) ListCarousel(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.SUCCESS
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))

}
