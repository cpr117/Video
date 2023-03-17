// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Comment struct{}

// 添加删除评论
func (*Comment) ReplyComment(c *gin.Context) {
	r.SendCode(c, commentService.SaveComment(utils.BindJson[req.SaveComment](c)))
}

// 软删除
func (*Comment) DeleteComment(c *gin.Context) {
	r.SendCode(c, commentService.SoftDeleteComment(utils.BindQuery[req.DeleteComment](c)))
}

// 获取前台列表
func (*Comment) GetFrontList(c *gin.Context) {
	r.SuccessData(c, commentService.GetFrontList(utils.BindQuery[req.GetFrontComments](c)))
}
