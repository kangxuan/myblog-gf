package common

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/internal/consts"
)

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

// GetTagByTagId 根据tag_id获取标签
func (t *tagCommonService) GetTagByTagId(tagId int) *TagSelectFields {
	var (
		result       = new(TagSelectFields)
		ctx          = gctx.New()
		ex     int64 = 3600
	)
	v, err := g.Redis().Get(ctx, consts.TagByTagId+gconv.String(tagId))
	if err != nil {
		panic(err)
	}
	if v.IsStruct() {
		_ = v.Struct(&result)
	} else {
		_ = g.Model("tag").Where("tag_id = ?", tagId).Fields("tag_id, tag_name").Scan(&result)
		_, err = g.Redis().Set(ctx, consts.TagByTagId+gconv.String(tagId), result, gredis.SetOption{
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
