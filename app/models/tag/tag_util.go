package tag

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idstr string) (tag Tag) {
	database.DB.Where("id", idstr).First(&tag)
	return
}

func GetBy(field, value string) (tag Tag) {
	database.DB.Where("? = ?", field, value).First(&tag)
	return
}

func All() (tags []Tag) {
	database.DB.Scopes(ScopeEnable).Find(&tags)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Tag{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (tags []Tag, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Tag{}),
		&tags,
		app.V1URL(database.TableName(&Tag{})),
		perPage,
	)
	return
}
