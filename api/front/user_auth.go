// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

//// 根据Token获取用户信息
//func (*UserAuth) GetInfo(c *gin.Context) {
//	r.SuccessData(c, userService.GetUserInfo(utils.GetFromContext[int](c, "user_info_id")))
//}

// 用户登录
func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.UserName, loginReq.PassWord)
	r.SendData(c, code, loginVo)
}

// 退出登录
func (*UserAuth) Logout(c *gin.Context) {
	logoutReq := utils.BindValidJson[req.Logout](c)
	userService.Logout(c, logoutReq.UserName)
	r.Success(c)
}

// 注册用户
func (*UserAuth) Register(c *gin.Context) {
	r.SendCode(c, userService.Register(utils.BindValidJson[req.Register](c), 2))
}

// GetCaptcha 获取验证码
func (*UserAuth) GetCaptcha(c *gin.Context) {
	r.SendCode(c, userService.GetCaptcha(c, utils.BindValidJson[req.RegisterCode](c)))
}
