package article_tag

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (articleTag ArticleTag) {
	database.DB.Where("id", idstr).First(&articleTag)
	return
}

func GetBy(field, value string) (articleTag ArticleTag) {
	database.DB.Where("? = ?", field, value).First(&articleTag)
	return
}

func All() (articleTags []ArticleTag) {
	database.DB.Find(&articleTags)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ArticleTag{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (articleTags []ArticleTag, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(ArticleTag{}),
		&articleTags,
		app.V1URL(database.TableName(&ArticleTag{})),
		perPage,
	)
	return
}
