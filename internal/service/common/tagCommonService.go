package common

import "github.com/gogf/gf/v2/frame/g"

var TagCommonService tagCommonService

type tagCommonService struct {
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
