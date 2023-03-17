// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Message struct{}

// 查询列表
func (*Message) GetMessageList(c *gin.Context) {
	r.SuccessData(c, messageService.GetMessageByUserId(utils.BindJson[req.GetMessage](c)))
}

// 添加新记录
func (*Message) SendMessage(c *gin.Context) {
	r.SendCode(c, messageService.SendMsg(
		utils.GetFromContext[int](c, "user_info_id"),
		utils.BindValidJson[req.SendMsg](c)))
}
