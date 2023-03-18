package routes

import (
	"go_api_skeleton/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 路由组
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判定手机号是否存在
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}

}
