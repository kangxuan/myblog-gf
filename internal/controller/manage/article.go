package manage

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var ArticleController articleController

type articleController struct{}

func (a *articleController) Create(ctx context.Context, req *manage.CreateArticleReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.CreateArticleReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}
	res = manage2.ArticleService.Create(ctx, params)
	return
}

func (a *articleController) Update(ctx context.Context, req *manage.UpdateArticleReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.UpdateArticleReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}
	res = manage2.ArticleService.Update(ctx, params)
	return
}

func (a *articleController) Delete(ctx context.Context, req *manage.DeleteArticleReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.DeleteArticleReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}
	res = manage2.ArticleService.Delete(ctx, params)
	return
}

func (a *articleController) GetA(ctx context.Context, req *manage.GetAArticleReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.GetAArticleReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}
	res = manage2.ArticleService.GetAArticle(ctx, params)
	return
}

func (a *articleController) GetList(ctx context.Context, req *manage.GetArticleListReq) (res *api.CommonJsonRes, err error) {
	params := new(manage.GetArticleListReq)
	if err = gconv.Struct(req, &params); err != nil {
		return nil, err
	}
	res = manage2.ArticleService.GetArticleList(ctx, params)
	return
}
