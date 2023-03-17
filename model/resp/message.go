// @User CPR
package resp

type FrontMsgVo struct {
	MsgId     int    `json:"msg_id"`
	MsgType   int    `json:"msg_type"`
	SendId    int    `json:"send_id"`
	ReceiveId int    `json:"receive_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
