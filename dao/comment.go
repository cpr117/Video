// @User CPR
package dao

import (
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Comment struct{}

func (*Comment) GetFrontList(req req.GetFrontComments) (list []resp.FrontCommentVo, total int64) {
	DB.Select("u.nickname, u.avatar, c.id, c.user_id, c.content, c.created_time, c.like_count, c.dislike_count, c.reply_comment_id").
		Table("comment c").
		Joins("LEFT JOIN user_info u ON c.user_id = u.id").
		Where("c.video_type = ? AND c.video_id = ?", req.VideoType, req.VideoId).
		Where("c.reply_comment_id is NULL").
		Count(&total).
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Order("c.created_time DESC").
		Find(&list)
	return
}

func (*Comment) GetReplyList(ids []int) (list []resp.FrontCommentVo) {
	DB.Select("u.nickname, u.avatar, c.id, c.user_id, c.content, c.created_time, c.like_count, c.dislike_count, c.reply_comment_id").
		Table("comment c").
		Joins("LEFT JOIN user_info u ON c.user_id = u.id").
		Where("c.reply_root_id IN ?", ids).
		Order("c.created_time DESC").
		Find(&list)
	return
}
