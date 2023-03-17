// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils/r"
)

type Partition struct{}

func (*Partition) GetPartitionList() []resp.GetPartition {
	return partitionDao.GetPartitionList()
	//return []string{"动画", "番剧", "国创", "音乐", "舞蹈", "游戏", "科技", "数码", "生活", "鬼畜", "时尚", "娱乐", "影视", "纪录片", "电影", "电视剧"}
}

func (*Partition) AddPartition(req req.AddPartition) int {
	dao.Create(&model.Partition{Name: req.Name})
	return r.OK
}

func (*Partition) DeletePartition(id int) int {
	dao.Delete[model.Partition]("id = ?", id)
	return r.OK
}

func (*Partition) UpdatePartition(req req.UpdatePartition) int {
	dao.Create(&model.Partition{
		Universal: model.Universal{Id: req.Id},
		Name:      req.Name,
	})
	return r.OK
}
