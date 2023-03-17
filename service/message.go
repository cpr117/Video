// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
)

type Message struct{}

func (*Message) GetMessageByUserId(req req.GetMessage) (ret resp.PageResult[[]resp.FrontMsgVo]) {
	ret.List, ret.Total = messageDao.GetMessageByUserId(req)
	ret.PageNum = req.PageNum
	ret.PageSize = req.PageSize
	return
}

func (*Message) SendMsg(sendId int, req req.SendMsg) (ret int) {
	msg := utils.CopyProperties[model.Message](ret) // vo->po
	msg.SendId = sendId
	dao.Create(&msg)
	return r.OK
}
