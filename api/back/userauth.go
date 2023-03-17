// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

// 根据Token获取用户信息
func (*UserAuth) GetInfo(c *gin.Context) {
	r.SuccessData(c, userService.GetInfo(utils.GetFromContext[int](c, "user_info_id")))
}

// 后台用户登录
func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.UserName, loginReq.PassWord)
	r.SendData(c, code, loginVo)
}

// 后台用户登出
func (*UserAuth) Logout(c *gin.Context) {
	userService.Logout(c, utils.GetFromContext[string](c, "username"))
	r.Success(c)
}

// 后台用户注册
func (*UserAuth) Register(c *gin.Context) {
	userService.Register(utils.BindValidJson[req.Register](c), 1)
}
