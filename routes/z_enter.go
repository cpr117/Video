// @User CPR
package routes

import (
	"VideoWeb/api/back"
	"VideoWeb/api/front"
)

// 前台接口
var (
	fAnimeAPI    front.Anime    // 动漫
	fUserAPI     front.User     // 用户
	fUserAuthAPI front.UserAuth // 用户认证
	fCateAPI     front.Category // 分类
	//liveAPI     front.Live     // 直播
	fCommentAPI front.Comment // 评论
	fMessageAPI front.Message // 留言
	//fBiliInfoAPI   front.BiliInfo   // B站信息
	fVideoAPI      front.Video      // 视频
	fCollectionAPI front.Collection // 归档
	fPartitionAPI  front.Partition  // 分区
)

// 后台接口
var (
	userAuthAPI back.UserAuth // 用户认证
	userAPI     back.User     // 用户
	//biliInfoAPI  back.BiliInfo  // B站信息
	partitionAPI back.Partition // 分区
	adminAPI     back.Admin     // 管理员
	videoAPI     back.Video     // 视频
	commentAPI   back.Comment   // 评论
)
