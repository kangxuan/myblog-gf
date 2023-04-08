package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/internal/service/common"
	"myblog-gf/utility"
	"time"
)

var TagService tagService

type tagService struct{}

// Create 创建标签
func (t *tagService) Create(_ context.Context, req *manage.CreateTagReq) (res *api.CommonJsonRes, err error) {
	unix := time.Now().Unix()

	isExisted, err := tagNameExisted(req.TagName, 0)
	if err != nil {
		return
	}
	if isExisted {
		res = utility.CommonResponse.ErrorMsg("标签名称已存在")
		return
	}

	tagModel := g.Model("tag")
	_, err = tagModel.Data(g.Map{
		"tag_name":    req.TagName,
		"create_time": unix,
		"update_time": unix,
	}).Insert()
	if err != nil {
		return
	}
	res = utility.CommonResponse.SuccessMsg("创建成功", nil)
	return
}

// Update 修改标签
func (t *tagService) Update(_ context.Context, req *manage.UpdateTagReq) (res *api.CommonJsonRes, err error) {
	unix := time.Now().Unix()

	isExisted, err := tagNameExisted(req.TagName, req.TagId)
	if err != nil {
		return
	}
	if isExisted {
		res = utility.CommonResponse.ErrorMsg("标签名称已存在")
		return
	}

	tagModel := g.Model("tag")
	_, err = tagModel.Data(g.Map{
		"tag_name":    req.TagName,
		"update_time": unix,
	}).Where("tag_id = ?", req.TagId).Update()
	if err != nil {
		return
	}

	// 清空缓存
	common.TagCommonService.DeleteTagCache(req.TagId)
	res = utility.CommonResponse.SuccessMsg("修改标签成功", nil)
	return
}

// Delete 删除标签
func (t *tagService) Delete(_ context.Context, req *manage.DeleteTagReq) (res *api.CommonJsonRes, err error) {
	unix := time.Now().Unix()

	_, err = g.Model("tag").Where("tag_id = ?", req.TagId).Data(g.Map{
		"is_delete":   1,
		"update_time": unix,
	}).Update()
	if err != nil {
		return
	}

	// 清空缓存
	common.TagCommonService.DeleteTagCache(req.TagId)
	res = utility.CommonResponse.SuccessMsg("删除标签成功", nil)
	return
}

// GetATag 获取一个标签
func (t *tagService) GetATag(_ context.Context, req *manage.GetATagReq) (res *api.CommonJsonRes, err error) {
	r, err := g.Model("tag").Fields("tag_id", "tag_name").Where("tag_id = ?", req.TagId).One()
	if err != nil {
		return
	}
	res = utility.CommonResponse.SuccessMsg("获取单个标签成功", r)
	return
}

// GetAllTag 获取标签列表
func (t *tagService) GetAllTag(_ context.Context, req *manage.GetAllTagReq) (res *api.CommonJsonRes, err error) {
	model := g.Model("tag").Fields("tag_id", "tag_name").Where("is_delete = ?", 0)
	if req.TagId > 0 {
		model.Where("tag_id = ?", req.TagId)
	}
	if req.TagName != "" {
		model.Where("tag_name like %?%", req.TagName)
	}
	r, err := model.All()
	if err != nil {
		return
	}
	res = utility.CommonResponse.SuccessMsg("获取所有标签列表成功", r)
	return
}

// tagNameExisted 判断标签名称是否存在
func tagNameExisted(tagName string, tagId int) (isExisted bool, err error) {
	model := g.Model("tag").Where("tag_name = ?", tagName)
	if tagId > 0 {
		model.Where("tag_id = ", tagId)
	}
	r, err := model.One()
	if err != nil {
		return
	}

	if !r.IsEmpty() {
		isExisted = true
	}

	return
}
