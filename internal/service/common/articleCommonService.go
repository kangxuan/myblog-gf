package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

var ArticleCommonService articleCommonService

type articleCommonService struct{}

// ArticleIdExisted 判断article_id是否存在
func (a *articleCommonService) ArticleIdExisted(articleId int) bool {
	r, err := g.Model("article").Where("article_id = ?", articleId).One()
	if err != nil {
		panic(err)
	}

	if r.IsEmpty() {
		return false
	}
	return true
}
