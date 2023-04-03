package common

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/internal/consts"
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
	var (
		result         = new(CategorySelectFields)
		ctx            = gctx.New()
		ex       int64 = 3600
		cacheKey       = consts.CategoryById + gconv.String(categoryId)
	)
	// 先查询缓存
	v, err := g.Redis().Get(ctx, cacheKey)
	if err != nil {
		panic(err)
	}
	if v.IsStruct() {
		err = v.Struct(&result)
		if err != nil {
			panic(err)
		}
	} else {
		_ = g.Model("category").Where("category_id = ?", categoryId).Fields("category_id, category_name").Scan(&result)
		// 设置缓存
		_, err = g.Redis().Set(ctx, cacheKey, result, gredis.SetOption{
			TTLOption: gredis.TTLOption{
				EX: &ex,
			},
		})
		if err != nil {
			panic(err)
		}
	}

	return result
}

// DeleteCategoryCache 删除缓存
func (c *categoryCommonService) DeleteCategoryCache(categoryId int) {
	var (
		ctx      = gctx.New()
		cacheKey = consts.CategoryById + gconv.String(categoryId)
	)

	_, err := g.Redis().Del(ctx, cacheKey)
	if err != nil {
		panic(err)
	}
}
