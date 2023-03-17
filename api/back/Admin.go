// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (*Admin) GetList(c *gin.Context) {
	r.SuccessData(c, adminService.GetList(utils.BindJson[req.GetAdminList](c)))
}

func (*Admin) Create(c *gin.Context) {
	adminService.Create(utils.BindJson[req.CreateAdmin](c))
	r.Success(c)
}

func (*Admin) Update(c *gin.Context) {
	r.SuccessData(c, adminService.Update(utils.BindJson[req.UpdateAdmin](c)))

}
