package manage

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var (
	CategoryController = cCategory{}
)

type cCategory struct{}

func (c *cCategory) Create(ctx context.Context, req *manage.CreateCategoryReq) (res *api.CommonJsonRes, err error) {
	var reqParams *manage.CreateCategoryReq
	if err = gconv.Struct(req, &reqParams); err != nil {
		return nil, err
	}
	res = manage2.CategoryService.Create(ctx, reqParams)
	return
}

func (c *cCategory) Update(ctx context.Context, req *manage.UpdateCategoryReq) (res *api.CommonJsonRes, err error) {
	var reqParams *manage.UpdateCategoryReq
	if err = gconv.Struct(req, &reqParams); err != nil {
		return nil, err
	}
	res = manage2.CategoryService.Update(ctx, reqParams)
	return
}

func (c *cCategory) Delete(ctx context.Context, req *manage.DeleteCategoryReq) (res *api.CommonJsonRes, err error) {
	var reqParams *manage.DeleteCategoryReq
	if err = gconv.Struct(req, &reqParams); err != nil {
		return nil, err
	}
	res = manage2.CategoryService.Delete(ctx, reqParams)
	return
}
