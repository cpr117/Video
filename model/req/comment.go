// @User CPR
package req

// GetFrontComments
// 这个是列表
// 通过视频类型和视频id可以锁定到具体的一个视频
type GetFrontComments struct {
	PageQuery
	VideoType int `json:"video_type" form:"video_type"`
	VideoId   int `json:"video_id" form:"video_id"`
}

type SaveComment struct {
	UserId         int    `json:"user_id"`
	ReplyCommentId int    `json:"reply_comment_id"`
	ReplyRootId    int    `json:"reply_root_id"`
	VideoId        int    `json:"video_id"`
	VideoType      int    `json:"video_type"`
	Content        string `json:"content"`
}

type DeleteComment struct {
	UserId    int `json:"user_id"`
	CommentId int `json:"comment_id"`
}

type GetCommentList struct {
	PageQuery
	VideoType int `json:"video_type" form:"video_type"`
	VideoId   int `json:"video_id" form:"video_id"`
}
