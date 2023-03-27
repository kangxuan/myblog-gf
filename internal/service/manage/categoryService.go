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

func (c *sCategory) Create(_ context.Context, req *manage.CreateCategoryReq) (res *api.CommonJsonRes) {
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

func (c *sCategory) Update(_ context.Context, req *manage.UpdateCategoryReq) (res *api.CommonJsonRes) {
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

func (c *sCategory) Delete(_ context.Context, req *manage.DeleteCategoryReq) (res *api.CommonJsonRes) {
	currentTime := time.Now().Unix()
	category := g.Model("category").Safe()
	_, err := category.Data(g.Map{
		"is_delete":   1,
		"update_time": currentTime,
	}).Where(g.Map{
		"category_id": req.CategoryId,
	}).Update()
	if err != nil {
		return utility.CommonResponse.ErrorMsg("删除分类失败")
	}
	return utility.CommonResponse.SuccessMsg("删除成功", nil)
}

func (c *sCategory) GetA(_ context.Context, req *manage.GetACategoryReq) (res *api.CommonJsonRes) {
	category := g.Model("category").Safe()
	categoryItem, err := category.Fields("category_id, category_name, parent_id").Where(g.Map{"category_id": req.CategoryId}).One()
	if err != nil {
		return utility.CommonResponse.ErrorMsg("获取单个分类失败")
	}
	return utility.CommonResponse.SuccessMsg("获取成功", categoryItem)
}

func (c *sCategory) GetList(_ context.Context, req *manage.GetCategoryListReq) (res *api.CommonJsonRes) {
	category := g.Model("category").Safe()

	where := make(map[string]interface{})
	if req.ParentId != 0 {
		where["parent_id"] = req.ParentId
	}
	if req.CategoryType != 0 {
		where["category_type"] = req.CategoryType
	}
	if req.CategoryName != "" {
		where["category_name like"] = "%" + req.CategoryName + "%"
	}

	if len(where) != 0 {
		category = category.Where(where)
	}

	category1 := category.Fields("category_id", "category_name", "parent_id")

	list := utility.PagingList(category1, req.Page, req.PageSize)
	page := utility.Page(category, req.Page, req.PageSize)

	return utility.CommonResponse.SuccessMsg("获取列表成功", g.Map{
		"list": list,
		"page": page,
	})
}
