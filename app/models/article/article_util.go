package article

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"gorm.io/gorm"
)

// ScopeClientView 客户端查询范围
func ScopeClientView(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at is null").Where("status = 0")
}

func Get(idstr string) (article Article) {
	database.DB.Where("id", idstr).
		Preload("Category").
		Preload("ArticleTag").
		Preload("ArticleTag.Tag").
		Scopes(ScopeClientView).
		First(&article)

	// 指定字段更新
	database.DB.Model(&article).
		Select("pv").
		UpdateColumn("pv", gorm.Expr("pv + ?", 1))

	return
}

func GetPreNext(idstr string) (pre PreNextArticleStruct, next PreNextArticleStruct) {
	database.DB.Model(&Article{}).Where("id > ?", idstr).
		Scopes(ScopeClientView).
		First(&next)

	database.DB.Model(&Article{}).Where("id < ?", idstr).
		Scopes(ScopeClientView).
		First(&pre)
	return
}

func GetBy(field, value string) (article Article) {
	// 预加载
	database.DB.Where("? = ?", field, value).
		Preload("Category").
		Preload("ArticleTag").
		Preload("ArticleTag.Tag").
		Scopes(ScopeClientView).
		First(&article)
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

	// 查询分类
	if !helpers.Empty(c.Query("category_id")) {
		db.Where("category_id", c.Query("category_id"))
	}
	// 查询标签
	if !helpers.Empty(c.Query("tag_id")) {
		var whereTags = "find_in_set(" + c.Query("tag_id") + ",`tag_ids`)"
		db.Where(whereTags)
	}

	db.Scopes(ScopeClientView).Order("id desc")

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
