// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"VideoWeb/utils/upload"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Video struct{}

func (*Video) UploadVideo(c *gin.Context) int {
	video := utils.GetFormValue(c, "video")
	//fmt.Println(video)
	videoMsg := req.SaveVideo{}
	utils.Json.Unmarshal(utils.GetFormTest(c, "msg"), &videoMsg)
	if CheckVideoMsg(videoMsg) {
		oss := upload.NewOSS()
		filePath, _, uploadErr := oss.UploadFile(video)
		if uploadErr != nil {
			fmt.Println("uploadErr: ", uploadErr)
			return r.ERROR_FILE_UPLOAD
		}
		//if err := c.SaveUploadedFile(video, filePath); err != nil {
		//	fmt.Println("err: ", err)
		//	utils.Logger.Info("UploadVideo", zap.Error(err))
		//	return r.ERROR_FILE_UPLOAD
		//}
		CreateVideo(videoMsg, filePath)
	}
	return r.OK
}

func CheckVideoMsg(req req.SaveVideo) bool {
	if dao.Count[model.Partition]("id", req.PartitionId) == 0 {
		utils.Logger.Error("CheckVideoMsg 分区不存在")
		return false
	}
	if dao.Count[model.UserAuth]("id", req.AuthorId) == 0 {
		utils.Logger.Error("CheckVideoMsg 用户id不存在")
		return false
	}
	return true
}

//func (*Video)

func CreateVideo(req req.SaveVideo, filePath string) {
	fmt.Println(req)
	dao.Create(&model.Video{
		Title:       req.Title,
		Url:         filePath,
		Desc:        req.Desc,
		Cover:       "req.VideoCover", // TODO: 上传视频封面
		AuthorId:    req.AuthorId,
		Status:      0,
		PartitionId: req.PartitionId,
	})
}

func (*Video) GetVideoList(req req.GetVideos) (ret resp.PageResult[[]resp.VideoVo]) {
	ret.List, ret.Total = videoDao.GetVideoList(req)
	ret.PageNum = req.PageNum
	ret.PageSize = req.PageSize
	return
}

func (*Video) GetVideoById(VideoId int) (resp.VideoVo, int) {
	if VideoId == 0 {
		utils.Logger.Info("GetVideoById: VideoId is null")
		return resp.VideoVo{}, r.ERROR_VIDEOID_WRONG
	}
	video := utils.CopyProperties[resp.VideoVo](videoDao.GetVideoById(VideoId))
	//获取视频资源
	return video, r.OK
}

// 前台的视频删除为软删除 即修改status为1
func (*Video) DeleteVideo(VideoId int) int {
	if VideoId == 0 {
		utils.Logger.Info("DeleteVideo: VideoId is null")
		return r.ERROR_VIDEOID_WRONG
	}
	videoDao.SoftDeleteVideo(VideoId)
	return r.OK
}

// 后台更新
func (*Video) UpdateVideo(req req.UpdateVideo) {
	data := utils.CopyProperties[model.Video](req)
	dao.Update[model.Video](&data)
}

func (*Video) HardDeleteVideo(VideoId int) {
	dao.Delete[model.Video]("id = ?", VideoId)
}
