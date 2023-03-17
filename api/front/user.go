// @User CPR
package front

import (
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) GetInfo(c *gin.Context) {
	r.SuccessData(c, userService.GetInfo(utils.GetFromContext[int](c, "user_info_id")))
}
