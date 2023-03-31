//Package article_tag 模型
package article_tag

import (
	"gohub/app/models"
	"gohub/app/models/tag"
	"gohub/pkg/database"
)

type ArticleTag struct {
	models.BaseModel

	// Put fields in here
	TagId     uint64 `json:"tag_id"`
	ArticleId uint64 `json:"article_id"`

	Tag tag.Tag `json:"tag"`
	models.CommonCreatedAtField
}

func (articleTag *ArticleTag) Create() {
	database.DB.Create(&articleTag)
}

func (articleTag *ArticleTag) Save() (rowsAffected int64) {
	result := database.DB.Save(&articleTag)
	return result.RowsAffected
}

func (articleTag *ArticleTag) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&articleTag)
	return result.RowsAffected
}
