package utility

import (
	"github.com/gogf/gf/v2/database/gdb"
)

// PagingList 获取分页数据
func PagingList(model *gdb.Model, page int, pageSize int) gdb.Result {
	//offset := (page - 1) * pageSize
	//result, err := model.Offset(offset).Limit(pageSize).All()
	result, err := model.Page(page, pageSize).All()
	if err != nil {
		panic(err)
	}
	return result
}

// Page 获取分页
func Page(model *gdb.Model, page int, pageSize int) map[string]int {
	count, err := model.Count()
	if err != nil {
		panic(err)
	}
	pageCount := count / pageSize
	if count%pageSize > 0 {
		pageCount++
	}
	if pageCount == 0 {
		pageCount = 1
	}

	return map[string]int{
		"page":       page,
		"page_total": pageCount,
		"count":      count,
	}
}
