package common

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api/v1/manage"
	"myblog-gf/internal/consts"
)

var ArticleCommonService articleCommonService

type articleCommonService struct{}

type ArticleListSelectFields struct {
	ArticleId        int    `json:"article_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Sort             int    `json:"sort"`
	CreateTime       int    `json:"create_time"`
	CategoryId       int    `json:"category_id"`
	TagId            int    `json:"tag_id"`
	CategoryName     string `json:"category_name"`
	TagName          string `json:"tag_name"`
	CreateTimeString string `json:"create_time_string"`
}

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

// GetArticleListMode 获取文章列表Mode
func (a *articleCommonService) GetArticleListMode(req *manage.ArticleListFields) *gdb.Model {
	model := g.Model("article a")
	model.LeftJoin("category_relation r1", "a.article_id = r1.relate_id and r1.type = "+gconv.String(consts.CategoryRelationTypeArticle)+" and r1.is_delete = 0")
	model.LeftJoin("tag_relation r2", "a.article_id = r2.relate_id and r2.type = "+gconv.String(consts.TagRelationTypeArticle)+" and r2.is_delete = 0")
	model.Where("a.is_delete = ?", 0)

	if req.ArticleId > 0 {
		model.Where("a.article_id = ?", req.ArticleId)
	}
	if len(req.Search) > 0 {
		model.Where("a.title like ? or a.content like ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	if req.CategoryId > 0 {
		model.Where("r1.category_id = ?", req.CategoryId)
	}
	if req.TagId > 0 {
		model.Where("r2.tag_id = ?", req.TagId)
	}
	return model
}
