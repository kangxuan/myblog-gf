// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	TagId      int         `json:"tagId"      description:""`
	TagName    string      `json:"tagName"    description:"标签名称"`
	IsDelete   uint        `json:"isDelete"   description:"删除 0-否 1-是"`
	CreateTime uint        `json:"createTime" description:"创建时间"`
	UpdateTime uint        `json:"updateTime" description:"更新时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   description:"变更时间"`
}
