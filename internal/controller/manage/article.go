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
