// @User CPR
package model

type Comment struct {
	Universal
	UserId         int    `gorm:"comment:评论用户id" json:"user_id"`
	ReplyCommentId int    `gorm:"comment:回复评论id" json:"reply_comment_id"`
	ReplyRootId    int    `gorm:"comment:回复根评论id" json:"reply_root_id"`
	VideoId        int    `gorm:"comment:评论主题id" json:"video_id"`
	VideoType      int    `gorm:"comment:评论主题类型" json:"video_type"`
	Content        string `gorm:"type:varchar(500);not null;comment:评论内容" json:"content"`
	IsDelete       int8   `gorm:"type:tinyint(1);not null;default:0;comment:是否删除(0.否 1.是)" json:"is_delete"`
	IsReview       int8   `gorm:"type:tinyint(1);not null;default:0;comment:是否审核(0.否 1.是)" json:"is_review"`
}
