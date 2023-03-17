// @User CPR
package dao

import (
	"MyBilibili/model/req"
	"MyBilibili/model/resp"
)

type Archive struct{}

func (*Archive) GetCollectionListByUserId(req req.GetCollection) (ret []resp.VideoVo, total int64) {
	DB.Model("video v").
		Joins("left join video_collection vc on v.id = vc.video_id").
		Joins("left join user_collection uc on uc.user_id = vc.user_id").
		Select("v.id, v.title, v.partition_id, v.video_desc, v.cover, v.author_id, v.video_url").
		Where("uc.user_id = ?", req.UserId).
		Where("uc.id = ?", req.CollectionId).
		Count(&total).
		Limit(req.PageSize).Offset((req.PageSize - 1) * req.PageSize).
		Find(&ret)
	return
}
