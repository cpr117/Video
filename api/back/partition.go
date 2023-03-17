// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Partition struct{}

func (*Partition) AddPartition(c *gin.Context) {
	r.SuccessData(c, partitionService.AddPartition(utils.BindJson[req.AddPartition](c)))
}

func (*Partition) DeletePartition(c *gin.Context) {
	r.SuccessData(c, partitionService.DeletePartition(utils.GetIntParam(c, "partition_id")))
}

func (*Partition) UpdatePartition(c *gin.Context) {
	partitionService.UpdatePartition(utils.BindJson[req.UpdatePartition](c))
	r.Success(c)
}

func (*Partition) GetPartitionList(c *gin.Context) {
	r.SuccessData(c, partitionService.GetPartitionList())
}
