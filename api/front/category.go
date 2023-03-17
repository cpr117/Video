// @User CPR
package front

import (
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Category struct{}

func (*Category) GetCategoryList(c *gin.Context) {
	r.SuccessData(c, categoryService.GetFrontList())
}
