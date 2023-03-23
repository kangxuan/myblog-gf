// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table category for DAO operations like Where/Data.
type Category struct {
	g.Meta       `orm:"table:category, do:true"`
	CategoryId   interface{} //
	CategoryType interface{} // 分类类型
	CategoryName interface{} // 分类名称
	ParentId     interface{} // 上级ID
	IsDelete     interface{} // 删除 0-否 1-是
	CreateTime   interface{} // 创建时间
	UpdateTime   interface{} // 更新时间
	UpdateAt     *gtime.Time // 变更时间
}
