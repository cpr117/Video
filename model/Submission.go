// @User CPR
package model

type Submission struct {
	Universal
	SubType int8 `gorm:"type:tinyint(1);not null;comment:提交类型" json:"sub_type"` // 0-视频 1-专栏 2-音频 3-贴纸 4-视频素材 todo:5-互动视频

}
