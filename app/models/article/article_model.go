//Package article 模型
package article

import (
	"gohub/app/models"
	"gohub/app/models/article_tag"
	"gohub/app/models/category"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
)

type Article struct {
	models.BaseModel

	// Put fields in here
	Title       string            `json:"title,omitempty"`
	CategoryId  uint64            `json:"category_id"`
	Status      uint64            `json:"status"`
	Content     string            `json:"content"`
	HtmlContent string            `json:"html_content"`
	Cover       string            `json:"cover"`
	SubMessage  string            `json:"sub_message"`
	Pv          uint64            `json:"pv"`
	IsEncrypt   uint64            `json:"is_encrypt,omitempty"`
	PublishAt   helpers.LocalTime `json:"publish_at,omitempty"`
	DeletedAt   helpers.LocalTime `json:"deleted_at,omitempty"`

	// 通过 category_id 关联分类
	Category   category.Category        `json:"category"`
	ArticleTag []article_tag.ArticleTag `json:"article_tags"`

	models.CommonTimestampsField
}

func (article *Article) Create() {
	database.DB.Create(&article)
}

func (article *Article) Save() (rowsAffected int64) {
	result := database.DB.Save(&article)
	return result.RowsAffected
}

func (article *Article) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&article)
	return result.RowsAffected
}
