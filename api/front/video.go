// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Video struct{}

// 视频上传
func (*Video) UploadVideo(c *gin.Context) {
	r.SendCode(c, videoService.UploadVideo(c))
}

// 获取视频列表
func (*Video) GetVideoList(c *gin.Context) {
	r.SuccessData(c, videoService.GetVideoList(utils.BindQuery[req.GetVideos](c)))
}

// 视频播放
func (*Video) GetVideoById(c *gin.Context) {
	data, code := videoService.GetVideoById(utils.GetIntParam(c, "id"))
	r.SendData(c, code, data)
}

// 视频删除
func (*Video) DeleteVideo(c *gin.Context) {
	r.SendCode(c, videoService.DeleteVideo(utils.GetIntParam(c, "id")))
}
