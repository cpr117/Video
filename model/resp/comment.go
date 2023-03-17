// @User CPR
package resp

import "time"

// 前台评论
type FrontCommentVo struct {
	ID             int              `json:"id"`
	UserId         int              `json:"user_id"`
	NickName       string           `json:"nickname"`
	Avatar         string           `json:"avatar"`
	Content        string           `json:"content"`
	CreatedAt      time.Time        `json:"reply_time"`
	LikeCount      int              `json:"like_count"`
	DislikeCount   int              `json:"dislike_count"`
	ReplyCommentId int              `json:"reply_comment_id"` //
	ReplyVoList    []FrontCommentVo `json:"reply_vo_list" gorm:"-"`
}

//// 前台回复
//type ReplyVo struct {
//	ID             int       `json:"id"`
//	CreatedAt      time.Time `json:"reply_time"`
//	UserId         int       `json:"user_id"`
//	NickName       string    `json:"nickname"`
//	Avatar         string    `json:"avatar"`
//	Content        string    `json:"content"`
//	ReplyUserId    int       `json:"reply_user_id"`
//	ReplyCommentId int       `json:"reply_comment_id"`
//	LikeCount      int       `json:"like_count"`
//	DislikeCount   int       `json:"dislike_count"`
//	ReplyVoList    []ReplyVo `json:"reply_vo_list" gorm:"-"`
//}
