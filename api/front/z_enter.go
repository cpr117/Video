// @User CPR
package front

import "VideoWeb/service"

var (
	userService     service.User     // 用户服务
	animeService    service.Anime    // 动漫服务
	categoryService service.Category // 动漫分类
	//liveService     service.Live     // 直播
	commentService    service.Comment    // 评论服务
	messageService    service.Message    // 留言服务 (这里是指私信)
	biliInfoService   service.BiliInfo   // B站信息服务
	videoService      service.Video      // 视频服务
	collectionService service.Collection // 归档服务
	partitionService  service.Partition  // 分区服务
)
