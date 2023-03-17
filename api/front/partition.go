// @User CPR
package front

import (
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Partition struct {
}

func (*Partition) GetPartitionList(c *gin.Context) {
	r.SuccessData(c, partitionService.GetPartitionList())
}
