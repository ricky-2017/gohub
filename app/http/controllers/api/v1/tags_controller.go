package v1

import (
	"gohub/app/models/tag"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	BaseAPIController
}

func (ctrl *TagsController) All(c *gin.Context) {
	tags := tag.All()
	response.Data(c, tags)
}

//func (ctrl *TagsController) Show(c *gin.Context) {
//    tagModel := tag.Get(c.Param("id"))
//    if tagModel.ID == 0 {
//        response.Abort404(c)
//        return
//    }
//    response.Data(c, tagModel)
//}
//
//func (ctrl *TagsController) Store(c *gin.Context) {
//
//    request := requests.TagRequest{}
//    if ok := requests.Validate(c, &request, requests.TagSave); !ok {
//        return
//    }
//
//    tagModel := tag.Tag{
//        FieldName:      request.FieldName,
//    }
//    tagModel.Create()
//    if tagModel.ID > 0 {
//        response.Created(c, tagModel)
//    } else {
//        response.Abort500(c, "创建失败，请稍后尝试~")
//    }
//}
//
//func (ctrl *TagsController) Update(c *gin.Context) {
//
//    tagModel := tag.Get(c.Param("id"))
//    if tagModel.ID == 0 {
//        response.Abort404(c)
//        return
//    }
//
//    if ok := policies.CanModifyTag(c, tagModel); !ok {
//        response.Abort403(c)
//        return
//    }
//
//    request := requests.TagRequest{}
//    bindOk, errs := requests.Validate(c, &request, requests.TagSave)
//    if !bindOk {
//        return
//    }
//    if len(errs) > 0 {
//        response.ValidationError(c, 20101, errs)
//        return
//    }
//
//    tagModel.FieldName = request.FieldName
//    rowsAffected := tagModel.Save()
//    if rowsAffected > 0 {
//        response.Data(c, tagModel)
//    } else {
//        response.Abort500(c, "更新失败，请稍后尝试~")
//    }
//}
//
//func (ctrl *TagsController) Delete(c *gin.Context) {
//
//    tagModel := tag.Get(c.Param("id"))
//    if tagModel.ID == 0 {
//        response.Abort404(c)
//        return
//    }
//
//    if ok := policies.CanModifyTag(c, tagModel); !ok {
//        response.Abort403(c)
//        return
//    }
//
//    rowsAffected := tagModel.Delete()
//    if rowsAffected > 0 {
//        response.Success(c)
//        return
//    }
//
//    response.Abort500(c, "删除失败，请稍后尝试~")
//}
