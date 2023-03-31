package common

import "github.com/gogf/gf/v2/frame/g"

var TagCommonService tagCommonService

type tagCommonService struct {
}

type TagSelectFields struct {
	TagId   int
	TagName string
}

// TagIdExisted 判断tag_id是否存在
func (t *tagCommonService) TagIdExisted(tagId int) bool {
	r, err := g.Model("tag").Where("tag_id = ? and is_delete = ?", tagId, 0).One()
	if err != nil {
		panic(err)
	}
	if r.IsEmpty() {
		return false
	}
	return true
}

func (t *tagCommonService) GetTagByTagId(tagId int) *TagSelectFields {
	result := new(TagSelectFields)
	_ = g.Model("tag").Where("tag_id = ?", tagId).Fields("tag_id, tag_name").Scan(&result)

	return result
}
