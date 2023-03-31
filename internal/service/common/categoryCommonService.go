package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

var CategoryCommonService categoryCommonService

type categoryCommonService struct{}

type CategorySelectFields struct {
	CategoryId   int
	CategoryName string
}

// CategoryIdExisted categoryId判断存在
func (c *categoryCommonService) CategoryIdExisted(categoryId int) bool {
	r, err := g.Model("category").Where("category_id = ? and is_delete = ?", categoryId, 0).One()
	if err != nil {
		panic(err)
	}
	if r.IsEmpty() {
		return false
	}
	return true
}

// GetCategoryByCategoryId 根据category_id获取分类
func (c *categoryCommonService) GetCategoryByCategoryId(categoryId int) *CategorySelectFields {
	result := new(CategorySelectFields)
	_ = g.Model("category").Where("category_id = ?", categoryId).Fields("category_id, category_name").Scan(&result)

	return result
}
