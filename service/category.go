package service

import (
	"context"
	"mall/dao"
	"mall/pkg/e"
	"mall/pkg/util"
	"mall/serializer"
)

type ListCategory struct {
}

func (service *ListCategory) List(ctx context.Context) serializer.Response {
	categoryDB := dao.NewCategoryDao(ctx)
	code := e.SUCCESS
	category, err := categoryDB.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))

}
