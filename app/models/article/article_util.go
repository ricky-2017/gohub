package article

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (article Article) {
	database.DB.Where("id", idstr).First(&article)
	return
}

func GetBy(field, value string) (article Article) {
	database.DB.Where("? = ?", field, value).First(&article)
	return
}

func All() (articles []Article) {
	database.DB.Find(&articles)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Article{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (articles []Article, paging paginator.Paging) {
	// 预加载
	db := database.DB.Model(Article{})
	db.Preload("Category").Preload("ArticleTag").Preload("ArticleTag.Tag")

	if !helpers.Empty(c.Query("category_id")) {
		db.Where("category_id", c.Query("category_id"))
	}
	if !helpers.Empty(c.Query("tag_id")) {
		var whereTags = "find_in_set(" + c.Query("tag_id") + ",`tag_ids`)"
		db.Where(whereTags)
	}

	// @todo 结合scope 封装查询
	paging = paginator.Paginate(
		c,
		db,
		&articles,
		app.V1URL(database.TableName(&Article{})),
		perPage,
	)
	return
}
