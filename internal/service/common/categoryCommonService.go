package common

import "github.com/gogf/gf/v2/frame/g"

var CategoryCommonService categoryCommonService

type categoryCommonService struct{}

// CategoryIdExisted categoryId判断存在
func (c *categoryCommonService) CategoryIdExisted(categoryId int) bool {
	r, err := g.Model("tag").Where("category_id = ? and is_delete = ?", categoryId, 0).One()
	if err != nil {
		panic(err)
	}
	if r.IsEmpty() {
		return false
	}
	return true
}
