package category

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (category Category) {
	database.DB.Where("id", idstr).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ?", field, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Scopes(ScopeEnable).Select([]string{"id", "name"}).Order("id desc").Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	query := database.DB.Model(Category{})
	paging = paginator.Paginate(
		c,
		query,
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
