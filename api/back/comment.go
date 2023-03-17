// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Comment struct{}

// 硬删除
func (*Comment) DeleteComment(c *gin.Context) {
	r.SendCode(c, commentService.HardDeleteComment(utils.BindQuery[req.DeleteComment](c)))
}

// 获取某个视频下的所有评论
// 后台和前端获取方法差不多 所以暂时是直接用了 后续可以进行修改
func (*Comment) GetListByVideoId(c *gin.Context) {
	r.SuccessData(c, commentService.GetFrontList(utils.BindQuery[req.GetFrontComments](c)))
}
