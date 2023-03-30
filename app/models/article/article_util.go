package article

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
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
	paging = paginator.Paginate(
		c,
		database.DB.Model(Article{}),
		&articles,
		app.V1URL(database.TableName(&Article{})),
		perPage,
	)
	return
}
