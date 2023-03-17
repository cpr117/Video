// @User CPR
package back

import "VideoWeb/service"

var (
	userService      service.User      // 用户服务
	userAuthService  service.UserAuth  // 用户认证服务
	biliInfoService  service.BiliInfo  // B站信息服务
	commentService   service.Comment   // 评论服务
	adminService     service.Admin     // 管理员服务
	partitionService service.Partition // 分区服务
	videoService     service.Video     // 视频服务
)
