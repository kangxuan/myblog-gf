package manage

import "github.com/gogf/gf/v2/frame/g"

type CreateTagReq struct {
	g.Meta `path:"/tag" method:"post" tags:"标签" summary:"创建标签"`
	TagFields
}

type UpdateTagReq struct {
	g.Meta `path:"/tag/{tag_id}" method:"put" tags:"标签" summary:"更新标签"`
	TagId  int `p:"tag_id" dc:"标签ID" v:"required|min:1#标签ID必传|标签ID必须大于0"`
	TagFields
}

type DeleteTagReq struct {
	g.Meta `path:"/tag/{tag_id}" method:"delete" tags:"标签" summary:"删除标签"`
	TagId  int `p:"tag_id" dc:"标签ID" v:"required|min:1#标签ID必传|标签ID必须大于0"`
}

type GetATagReq struct {
	g.Meta `path:"/tag/{tag_id}" method:"get" tags:"标签" summary:"获取单个标签"`
	TagId  int `p:"tag_id" dc:"标签ID" v:"required|min:1#标签ID必传|标签ID必须大于0"`
}

type GetAllTagReq struct {
	g.Meta  `path:"/tag" method:"get" tags:"标签" summary:"获取所有标签"`
	TagId   int    `p:"tag_id" dc:"标签ID"`
	TagName string `p:"tag_name" dc:"标签名称"`
}

type TagFields struct {
	TagName string `json:"tag_name" dc:"标签名称" v:"required|max-length:10#标签名称必填|标签名称最大长度不能超过10个字"`
}
