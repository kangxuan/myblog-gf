package manage

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/internal/consts"
	"myblog-gf/internal/service/common"
	"myblog-gf/utility"
	"time"
)

var ArticleService articleService

type articleService struct{}

// Create 创建文章
func (a *articleService) Create(ctx context.Context, req *manage.CreateArticleReq) (res *api.CommonJsonRes) {
	if common.CategoryCommonService.CategoryIdExisted(req.CategoryId) {
		return utility.CommonResponse.ErrorMsg("分类ID不存在")
	}
	if req.TagId > 0 && common.TagCommonService.TagIdExisted(req.TagId) {
		return utility.CommonResponse.ErrorMsg("标签ID不存在")
	}

	unix := time.Now().Unix()
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 文章
		r, err := tx.Model("article").Data(g.Map{
			"title":       req.Title,
			"content":     req.Content,
			"sort":        req.Sort,
			"create_time": unix,
			"update_time": unix,
		}).Insert()
		if err != nil {
			return err
		}

		articleId, err := r.LastInsertId()
		if err != nil {
			return err
		}

		// 文章分类
		_, err = tx.Model("category_relation").Data(g.Map{
			"category_id": req.CategoryId,
			"type":        consts.CategoryRelationTypeArticle,
			"relate_id":   articleId,
			"create_time": unix,
			"update_time": unix,
		}).Insert()
		if err != nil {
			return err
		}

		// 文章标签
		if req.TagId > 0 {
			_, err = tx.Model("tag_relation").Data(g.Map{
				"tag_id":      req.TagId,
				"type":        consts.TagRelationTypeArticle,
				"relate_id":   articleId,
				"create_time": unix,
				"update_time": unix,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return utility.CommonResponse.ErrorMsg("创建文章失败")
	}
	return utility.CommonResponse.SuccessMsg("创建文章成功", nil)
}

// Update 更新文章
func (a *articleService) Update(ctx context.Context, req *manage.UpdateArticleReq) (res *api.CommonJsonRes) {
	if !common.ArticleCommonService.ArticleIdExisted(req.ArticleId) {
		return utility.CommonResponse.ErrorMsg("文章ID不存在")
	}

	if common.CategoryCommonService.CategoryIdExisted(req.CategoryId) {
		return utility.CommonResponse.ErrorMsg("分类ID不存在")
	}
	if req.TagId > 0 && common.TagCommonService.TagIdExisted(req.TagId) {
		return utility.CommonResponse.ErrorMsg("标签ID不存在")
	}

	unix := gtime.Timestamp
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 修改文章
		_, err := tx.Model("article").Where("article_id = ?", req.ArticleId).Data(g.Map{
			"title":       req.Title,
			"content":     req.Content,
			"sort":        req.Sort,
			"update_time": unix,
		}).Update()
		if err != nil {
			return err
		}

		// 删除分类关联并添加新关联
		_, err = tx.Model("category_relation").Where(g.Map{
			"relate_id": req.ArticleId,
			"type":      consts.CategoryRelationTypeArticle,
		}).Data(g.Map{"is_delete": 1}).Update()
		if err != nil {
			return err
		}
		_, err = tx.Model("category_relation").Data(g.Map{
			"category_id": req.CategoryId,
			"type":        consts.CategoryRelationTypeArticle,
			"relate_id":   req.ArticleId,
			"create_time": unix,
			"update_time": unix,
		}).Insert()
		if err != nil {
			return err
		}

		// 文章标签
		if req.TagId > 0 {
			_, err = tx.Model("tag_relation").Where(g.Map{
				"relate_id": req.ArticleId,
				"type":      consts.TagRelationTypeArticle,
			}).Data(g.Map{"is_delete": 1}).Update()
			if err != nil {
				return err
			}
			_, err = tx.Model("tag_relation").Data(g.Map{
				"tag_id":      req.TagId,
				"type":        consts.TagRelationTypeArticle,
				"relate_id":   req.ArticleId,
				"create_time": unix,
				"update_time": unix,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return utility.CommonResponse.ErrorMsg("修改文章失败")
	}
	return utility.CommonResponse.SuccessMsg("修改文章成功", nil)
}

// Delete 删除文章
func (a *articleService) Delete(ctx context.Context, req *manage.DeleteArticleReq) (res *api.CommonJsonRes) {
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除文章
		_, err := tx.Model("article").Where("article_id = ?", req.ArticleId).Update("is_delete = 1")
		if err != nil {
			return err
		}
		// 删除文章分类
		_, err = tx.Model("category_relation").Where(g.Map{
			"relate_id": req.ArticleId,
			"type":      consts.CategoryRelationTypeArticle,
		}).Update("is_delete = 1")
		if err != nil {
			return err
		}
		// 删除文章标签
		_, err = tx.Model("tag_relation").Where(g.Map{
			"relate_id": req.ArticleId,
			"type":      consts.TagRelationTypeArticle,
		}).Update("is_delete = 1")
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return utility.CommonResponse.ErrorMsg("删除文章失败")
	}
	return utility.CommonResponse.SuccessMsg("删除文章成功", nil)
}

// GetAArticle 获取单个文章
func (a *articleService) GetAArticle(_ context.Context, req *manage.GetAArticleReq) (res *api.CommonJsonRes) {
	model := g.Model("article a")
	model.LeftJoin("category_relation r1", "a.article_id = r1.relate_id and r1.type = "+gconv.String(consts.CategoryRelationTypeArticle)+" and r1.is_delete = 0")
	model.LeftJoin("tag_relation r2", "a.article_id = r2.relate_id and r2.type = "+gconv.String(consts.TagRelationTypeArticle)+" and r2.is_delete = 0")
	model.Where("a.article_id = ? and a.is_delete = ?", req.ArticleId, 0)
	model.Fields("a.article_id, a.title, a.content, a.sort, a.create_time, r1.category_id, r2.tag_id")
	r, err := model.One()
	if err != nil {
		panic(err)
	}

	article := r.Map()
	if len(article) > 0 {
		// 时间格式转换
		article["create_time"] = gtime.New(article["create_time"], "Y-m-d H:i:s")
		category := common.CategoryCommonService.GetCategoryByCategoryId(gconv.Int(article["category_id"]))
		article["category_name"] = category.CategoryName
		tag := common.TagCommonService.GetTagByTagId(gconv.Int(article["tag_id"]))
		article["tag_name"] = tag.TagName
	}
	return utility.CommonResponse.SuccessMsg("获取文章成功", article)
}
