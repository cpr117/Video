// @User CPR
package back

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Video struct{}

// 获取视频列表
func (*Video) GetVideoList(c *gin.Context) {
	r.SuccessData(c, videoService.GetVideoList(utils.BindQuery[req.GetVideos](c)))
}

func (*Video) UpdateVideo(c *gin.Context) {
	videoService.UpdateVideo(utils.BindJson[req.UpdateVideo](c))
	r.Success(c)
}

func (*Video) DeleteVideo(c *gin.Context) {
	videoService.HardDeleteVideo(utils.GetIntParam(c, "videoId"))
	r.Success(c)
}
