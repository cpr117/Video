// @User CPR
package model

type Favor struct {
	Universal
	VideoId        int    `gorm:"type:int;not null;comment:视频id" json:"video_id"`
	Avatar         string `gorm:"type:varchar(255);not null;comment:视频封面" json:"avatar"`
	Title          string `gorm:"type:varchar(255);not null;comment:视频标题" json:"title"`
	Classification string `gorm:"type:varchar(255);not null;comment:视频分类" json:"classification"`
}
