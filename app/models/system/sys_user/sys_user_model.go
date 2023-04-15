package sys_user

import (
	"gohub/app/models"
	"gohub/pkg/hash"
)

type SysUser struct {
	models.BaseModel

	UserAvatar             string `json:"user_avatar"`
	UserLastLoginIp        string `json:"user_last_login_ip"`
	UserName               string `json:"user_name"`
	UserNickname           string `json:"user_nickname"`
	UserPassword           string `json:"user_password"`
	UserPhone              string `json:"user_phone"`
	UserStatus             string `json:"user_status"`
	UserTokenVersion       string `json:"user_token_version"`
	UserGroup              string `json:"user_group"`
	LastUpdatePasswordTime string `json:"last_update_password_time"`

	models.CommonTimestampsField
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysUser) TableName() string {
	return "sys_user"
}

// ComparePassword 验证密码
func (sysUserModel *SysUser) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, sysUserModel.UserPassword)
}
