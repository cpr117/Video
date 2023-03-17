// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) GetList(c *gin.Context) {
	r.SuccessData(c, userService.GetList(utils.BindJson[req.GetUserList](c)))
}

func (*User) UpdateDisable(c *gin.Context) {
	userService.UpdateDisable(utils.BindJson[req.UpdateDisable](c))
	r.Success(c)
}

func (*User) DeleteUser(c *gin.Context) {
	userService.DeleteUser(utils.GetIntParam(c, "id"))
	r.Success(c)
}

func (*User) UpdateRole(c *gin.Context) {
	userService.UpdateRole(utils.BindJson[req.UpdateRole](c))
	r.Success(c)
}
