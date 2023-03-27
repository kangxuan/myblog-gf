package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"myblog-gf/api"
)

type CreateCategoryReq struct {
	g.Meta `path:"/category" method:"post" tags:"分类" summary:"创建分类"`
	CategoryFields
}

type UpdateCategoryReq struct {
	g.Meta     `path:"/category" method:"put" tags:"分类" summary:"更新分类"`
	CategoryId int `p:"id" dc:"分类ID" v:"required|min:1#分类ID必传|分类ID必须大于0"`
	CategoryFields
}

type DeleteCategoryReq struct {
	g.Meta     `path:"category" method:"delete" tags:"分类" summary:"删除分类"`
	CategoryId int `p:"id" dc:"分类ID" v:"required|min:1#分类ID必传|分类ID必须大于0"`
}

type GetACategoryReq struct {
	g.Meta     `path:"category/{id}" method:"GET" tags:"分类" summary:"获取分类"`
	CategoryId int `p:"id" dc:"分类ID" v:"required|min:1#分类ID必传|分类ID必须大于0"`
}

type GetCategoryListReq struct {
	g.Meta `path:"category" method:"GET" tags:"分类" summary:"分类列表"`
	api.PageParams
}

type CategoryFields struct {
	CategoryType int    `json:"category_type" dc:"分类类型" v:"required|in:1,2,3,4#分类类型必传|分类类型错误"`
	CategoryName string `json:"category_name" dc:"分类名称" v:"required|max-length:50#分类名称不能为空|分类名称最长不超过50个字"`
	ParentId     int    `json:"parent_id" dc:"上级分类ID" v:"required|min:0#上级分类必传|上级分类ID必须大于等于0"`
}
