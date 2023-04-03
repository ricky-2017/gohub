package v1

import (
	"gohub/app/models/article"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	BaseAPIController
}

func (ctrl *ArticlesController) Lists(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := article.Paginate(c, 10)
	response.Data(c, gin.H{
		"list":     data,
		"count":    pager.TotalCount,
		"page":     pager.CurrentPage,
		"pageSize": pager.PerPage,
	})
}

func (ctrl *ArticlesController) Show(c *gin.Context) {
	articleId := c.Param("id")
	articleModel := article.Get(articleId)
	if articleModel.ID == 0 {
		response.Abort404(c)
		return
	}
	pre, next := article.GetPreNext(articleId)

	response.Data(c, gin.H{
		"article": articleModel,
		"pre":     pre,
		"next":    next,
	})
}

//
//func (ctrl *ArticlesController) Store(c *gin.Context) {
//
//    request := requests.ArticleRequest{}
//    if ok := requests.Validate(c, &request, requests.ArticleSave); !ok {
//        return
//    }
//
//    articleModel := article.Article{
//        FieldName:      request.FieldName,
//    }
//    articleModel.Create()
//    if articleModel.ID > 0 {
//        response.Created(c, articleModel)
//    } else {
//        response.Abort500(c, "创建失败，请稍后尝试~")
//    }
//}
//
//func (ctrl *ArticlesController) Update(c *gin.Context) {
//
//    articleModel := article.Get(c.Param("id"))
//    if articleModel.ID == 0 {
//        response.Abort404(c)
//        return
//    }
//
//    if ok := policies.CanModifyArticle(c, articleModel); !ok {
//        response.Abort403(c)
//        return
//    }
//
//    request := requests.ArticleRequest{}
//    bindOk, errs := requests.Validate(c, &request, requests.ArticleSave)
//    if !bindOk {
//        return
//    }
//    if len(errs) > 0 {
//        response.ValidationError(c, 20101, errs)
//        return
//    }
//
//    articleModel.FieldName = request.FieldName
//    rowsAffected := articleModel.Save()
//    if rowsAffected > 0 {
//        response.Data(c, articleModel)
//    } else {
//        response.Abort500(c, "更新失败，请稍后尝试~")
//    }
//}
//
//func (ctrl *ArticlesController) Delete(c *gin.Context) {
//
//    articleModel := article.Get(c.Param("id"))
//    if articleModel.ID == 0 {
//        response.Abort404(c)
//        return
//    }
//
//    if ok := policies.CanModifyArticle(c, articleModel); !ok {
//        response.Abort403(c)
//        return
//    }
//
//    rowsAffected := articleModel.Delete()
//    if rowsAffected > 0 {
//        response.Success(c)
//        return
//    }
//
//    response.Abort500(c, "删除失败，请稍后尝试~")
//}
