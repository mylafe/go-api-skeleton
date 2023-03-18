// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "go_api_skeleton/app/http/controllers/api/v1"
	"go_api_skeleton/app/models/user"
	"go_api_skeleton/app/requests"
	"go_api_skeleton/pkg/response"

	"github.com/gin-gonic/gin"
)

// 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// 检测手机号是否存在
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
