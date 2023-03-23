// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryRelation is the golang structure for table category_relation.
type CategoryRelation struct {
	Id         int         `json:"id"         description:""`
	CategoryId uint        `json:"categoryId" description:"分类ID"`
	Type       uint        `json:"type"       description:"类型 1-文章"`
	RelateId   uint        `json:"relateId"   description:"关联ID"`
	IsDelete   uint        `json:"isDelete"   description:"删除 0-否 1-是"`
	CreateTime uint        `json:"createTime" description:"创建时间"`
	UpdateTime uint        `json:"updateTime" description:"更新时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   description:"变更时间"`
}
