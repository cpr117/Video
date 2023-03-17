// @User CPR
package model

import "reflect"

// 用户私信

type Message struct {
	Universal
	MsgType   int8   `gorm:"type:tinyint(1);not null;default:0;comment:消息类型(0-系统消息,1-用户私信)" json:"msg_type"`
	SendId    int    `gorm:"comment:发送者id" json:"send_id"`
	ReceiveId int    `gorm:"comment:接收者id" json:"receive_id"`
	Content   string `gorm:"type:varchar(255);comment:留言内容" json:"content"`
	// content 通常包含私信内容图片或者视频地址又或者是网页链接
}

func (c *Message) IsEmpty() bool {
	return reflect.DeepEqual(c, &Message{})
}
