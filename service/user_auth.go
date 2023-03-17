// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type UserAuth struct{}

func (*UserAuth) Logout(c *gin.Context) {
	// 登出更新数据库数据
	username := utils.GetFromContext[string](c, "username")
	userAuth := dao.GetOne[model.UserAuth]("username", username)
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{Id: userAuth.Id},
		LastLoginTime: time.Now(),
	})

	// 清除uuid信息
	uuid := utils.GetFromContext[string](c, "uuid")
	session := sessions.Default(c)
	session.Delete(KEY_USER + uuid)
	session.Save()
	utils.Redis.Del(KEY_USER + uuid)
}
