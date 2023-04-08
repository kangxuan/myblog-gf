package manage

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var TagController tagController

type tagController struct {
}

func (c *tagController) Create(ctx context.Context, req *manage.CreateTagReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.CreateTagReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}

	res, err = manage2.TagService.Create(ctx, params)
	return
}

func (c *tagController) Update(ctx context.Context, req *manage.UpdateTagReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.UpdateTagReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}

	res, err = manage2.TagService.Update(ctx, params)
	return
}

func (c *tagController) Delete(ctx context.Context, req *manage.DeleteTagReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.DeleteTagReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}

	res, err = manage2.TagService.Delete(ctx, params)
	return
}

func (c *tagController) GetATag(ctx context.Context, req *manage.GetATagReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.GetATagReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}

	res, err = manage2.TagService.GetATag(ctx, params)
	return
}

func (c *tagController) GetAllTag(ctx context.Context, req *manage.GetAllTagReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.GetAllTagReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}

	res, err = manage2.TagService.GetAllTag(ctx, params)
	return
}
