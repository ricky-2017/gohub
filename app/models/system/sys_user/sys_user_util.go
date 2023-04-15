package sys_user

import "gohub/pkg/database"

func GetUser(userName string) (sysUserModel SysUser) {
	database.DB.Where("user_name", userName).First(&sysUserModel)
	return
}
