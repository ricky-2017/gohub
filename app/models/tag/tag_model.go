//Package tag 模型
package tag

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gorm.io/gorm"
)

type Tag struct {
	models.BaseModel

	// Put fields in here
	Name   string `json:"name"`
	Status uint64 `json:"status"`

	models.CommonTimestampsField
}

func (tag *Tag) Create() {
	database.DB.Create(&tag)
}

func (tag *Tag) Save() (rowsAffected int64) {
	result := database.DB.Save(&tag)
	return result.RowsAffected
}

func (tag *Tag) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&tag)
	return result.RowsAffected
}

// ScopeEnable 已启用
func ScopeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("status = 0")
}
