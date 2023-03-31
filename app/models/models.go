package models

import (
	"github.com/spf13/cast"
	"gohub/pkg/helpers"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt helpers.LocalTime `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt helpers.LocalTime `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

type CommonCreatedAtField struct {
	CreatedAt helpers.LocalTime `gorm:"column:created_at;index;" json:"created_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
