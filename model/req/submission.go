// @User CPR
package req

type Submission struct {
	SubType int8 `json:"sub_type" validate:"required" label:"提交类型"` // 0-视频 1-专栏 2-音频 3-贴纸 4-视频素材 todo:5-互动视频
	VideoId int  `json:"video_id" validate:"required" label:"视频id"`
}
