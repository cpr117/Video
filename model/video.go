// @User CPR
package model

type Video struct {
	Universal
	Title       string `gorm:"type:varchar(255);not null;comment:视频名" json:"title"`
	Url         string `gorm:"type:varchar(255);not null;comment:视频地址" json:"url"`
	Cover       string `gorm:"type:varchar(255);not null;comment:视频封面" json:"cover"`
	Desc        string `gorm:"type:varchar(2000);not null;comment:视频描述" json:"desc"`
	PartitionId int    `gorm:"int(11);not null;comment:视频分类id" json:"partition_id"`
	Status      int8   `gorm:"type:tinyint(1);not null;comment:视频状态" json:"status"` // 0:未审核 1:审核通过 2:审核不通过
	AuthorId    int    `gorm:"int(11);comment:视频作者id" json:"author_id"`
}

type VideoResource struct {
	VideoId  int    `gorm:"type:int(11);not null;comment:视频id" json:"video_id"`
	Res360   string `gorm:"type:varchar(255);not null;comment:360P视频地址" json:"res_360"`
	Res480   string `gorm:"type:varchar(255);not null;comment:480P视频地址" json:"res_480"`
	Res720   string `gorm:"type:varchar(255);not null;comment:720P视频地址" json:"res_720"`
	Res1080  string `gorm:"type:varchar(255);not null;comment:1080P视频地址" json:"res_1080"`
	Original string `gorm:"type:varchar(255);not null;comment:原始视频地址" json:"original"`
	Duration string `gorm:"type:varchar(255);not null;comment:视频时长" json:"duration"`
}

type VideoCollection struct {
	VideoId      int `gorm:"type:int(11);not null;comment:视频id" json:"video_id"`
	CollectionId int `gorm:"type:int(11);not null;comment:收藏夹id" json:"collection_id"`
}

type VideoPartition struct {
	VideoId     int `gorm:"type:int(11);not null;comment:视频id" json:"video_id"`
	PartitionId int `gorm:"type:int(11);not null;comment:分区id" json:"partition_id"`
}
