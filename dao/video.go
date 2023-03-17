// @User CPR
package dao

import (
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"fmt"
)

type Video struct{}

func (*Video) GetVideoList(req req.GetVideos) (ret []resp.VideoVo, total int64) {
	//model.Video{}
	fmt.Println(req)
	db := DB.Table("video").
		Select("id", "title", "partition_id", "cover", "status", "desc", "author_id", "created_at", "updated_at")
	if req.VideoName != "" {
		db = db.Where("title like ?", "%"+req.VideoName+"%")
	}
	if req.VideoParId != 0 {
		db = db.Where("partition_id = ?", req.VideoParId)
	}
	db.Count(&total).
		Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).
		Find(&ret)
	fmt.Println(ret)
	return
}

func (*Video) GetVideoById(videoId int) (ret model.Video) {
	DB.Table("video").
		Where("id = ?", videoId).
		First(&ret)
	return
}

func (*Video) SoftDeleteVideo(videoId int) {
	DB.Table("video").
		Where("id = ?", videoId).
		Update("status", 1)
}
