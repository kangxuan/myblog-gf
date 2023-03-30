package manage

import (
	"context"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/internal/service/common"
	"myblog-gf/utility"
)

var ArticleService articleService

type articleService struct{}

func (a *articleService) Create(_ context.Context, req *manage.CreateArticleReq) (res *api.CommonJsonRes) {
	if common.CategoryCommonService.CategoryIdExisted(req.CategoryId) {
		return utility.CommonResponse.ErrorMsg("分类ID不存在")
	}
	if req.TagId > 0 && common.TagCommonService.TagIdExisted(req.TagId) {
		return utility.CommonResponse.ErrorMsg("标签ID不存在")
	}

	return
}
