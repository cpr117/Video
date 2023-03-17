// @User CPR
package back

import (
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type BiliInfo struct{}

func (*BiliInfo) GetInfo(c *gin.Context) {
	r.SuccessData(c, biliInfoService.GetFrontHomeInfo())
}
