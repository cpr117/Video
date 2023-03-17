// @User CPR
package req

type GetMessage struct {
	PageQuery
	SendId    int `json:"send_id" validate:"required"`
	ReceiveId int `json:"receive_id" validate:"required"`
	MsgType   int `json:"msg_type" validate:"required"`
}

type SendMsg struct {
	ReceiveId int    `json:"receive_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
	MsgType   int    `json:"msg_type" validate:"required"`
}
