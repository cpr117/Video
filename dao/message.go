// @User CPR
package dao

import (
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Message struct{}

func (*Message) GetMessageByUserId(req req.GetMessage) (ret []resp.FrontMsgVo, total int64) {
	DB.Table("message").
		Select("id", "send_id", "receive_id", "content", "created_at").
		Where("msg_type = ?", req.MsgType).
		Where("(send_id = ? AND receive_id = ?) OR (send_id = ? AND receive_id = ?)",
			req.SendId, req.ReceiveId, req.ReceiveId, req.SendId).
		Order("created_at DESC").
		Count(&total).
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Find(&ret)
	return ret, total
}
