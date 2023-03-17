// @User CPR
package dao

import (
	"VideoWeb/model/resp"
)

type Partition struct{}

func (*Partition) GetPartitionList() (ret []resp.GetPartition) {
	DB.Model("partition").
		Select("id, name").
		Find(&ret)
	return
}
