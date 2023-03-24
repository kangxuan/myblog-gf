package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/utility"
	"time"
)

var (
	CategoryService = sCategory{}
)

type sCategory struct {
}

func (s *sCategory) Create(_ context.Context, req *manage.CreateCategoryReq) (res *api.CommonJsonRes) {
	currentTime := time.Now().Unix()
	category := g.Model("category").Safe()
	_, err := category.Data(g.Map{
		"category_type": req.CategoryType,
		"category_name": req.CategoryName,
		"parent_id":     req.ParentId,
		"create_time":   currentTime,
		"update_time":   currentTime,
	}).Insert()
	if err != nil {
		return utility.CommonResponse.ErrorMsg("创建分类失败")
	}
	return utility.CommonResponse.SuccessMsg("创建成功", nil)
}

func (s *sCategory) Update(_ context.Context, req *manage.UpdateCategoryReq) (res *api.CommonJsonRes) {
	currentTime := time.Now().Unix()
	category := g.Model("category").Safe()
	_, err := category.Data(g.Map{
		"category_type": req.CategoryType,
		"category_name": req.CategoryName,
		"parent_id":     req.ParentId,
		"update_time":   currentTime,
	}).Where(g.Map{
		"category_id": req.CategoryId,
	}).Update()
	if err != nil {
		return utility.CommonResponse.ErrorMsg("更新分类失败")
	}
	return utility.CommonResponse.SuccessMsg("更新成功", nil)
}
