package manage

import "github.com/gogf/gf/v2/frame/g"

type CreateArticleReq struct {
	g.Meta `path:"/article" method:"POST" tags:"文章" summary:"创建文章"`
	ArticleFields
}

type UpdateArticleReq struct {
	g.Meta `path:"/article/{article_id}" method:"PUT" tags:"文章" summary:"修改文章"`
	ArticleIdFields
	ArticleFields
}

type DeleteArticleReq struct {
	g.Meta `path:"/article/{article_id}" method:"DELETE" tags:"文章" summary:"删除文章"`
	ArticleIdFields
}

type GetAArticleReq struct {
	g.Meta `path:"/article/{article_id}" method:"GET" tags:"文章" summary:"获取单个文章"`
	ArticleIdFields
}

type GetArticleListReq struct {
	g.Meta `path:"/article" method:"GET" tags:"文章" summary:"文章列表"`
	ArticleIdFields
	Search     string `p:"search" dc:"搜索内容"`
	CategoryId int    `p:"category_id" dc:"分类ID"`
	TagId      int    `p:"tag_id" dc:"标签ID"`
}

type ArticleIdFields struct {
	ArticleId int `p:"article_id" v:"required|min:1#文章ID必传|文章ID必须大于0" dc:"文章ID"`
}

type ArticleFields struct {
	Title      string `json:"title" v:"required|max-length:50#标题不能为空|标题最长不能超过50个字" dc:"标题"`
	Content    string `json:"content" v:"required#内容必传" default:"" dc:"内容"`
	Sort       int    `json:"sort" v:"min:0|max:999999#排序最小为0|排序最大为999999" default:"0" dc:"排序"`
	CategoryId int    `json:"category_id" v:"required|min:1#分类ID必传|分类ID必须大于0" dc:"分类ID"`
	TagId      int    `json:"tag_id" default:"0" dc:"标签ID"`
}
