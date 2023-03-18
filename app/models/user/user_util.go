// util
package user

import (
	"go_api_skeleton/pkg/database"
)

// 判断手机号是否存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
