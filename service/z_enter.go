// @User CPR
package service

import "VideoWeb/dao"

const (
	KEY_USER         = "user:"        // 记录用户
	KEY_DELETE       = "delete:"      //? 记录强制下线用户?
	KEY_ABOUT        = "about"        // 关于我信息
	KEY_BILI_CONFIG  = "bili_config"  // 网站配置信息
	KEY_VISITOR_AREA = "visitor_area" // 地域统计
	KEY_VIEW_COUNT   = "view_count"   // 访问数量

	KEY_UNIQUE_VISITOR_SET = "unique_visitor" // 唯一用户记录 set

	KEY_ANIME_USER_LIKE_SET = "anime_user_like:" // 动漫点赞 Set
	KEY_ANIME_LIKE_COUNT    = "anime_like_count" // 动漫点赞数
	KEY_ANIME_VIEW_COUNT    = "anime_view_count" // 动漫查看数

	KEY_COMMENT_LIKE_COUNT    = "comment_like_count"    // 评论点赞数
	KEY_COMMENT_DISLIKE_COUNT = "comment_dislike_count" // 评论点踩数

	KEY_PAGE = "page" // 页面封面
)

var (
	//videoDao dao.Video
	userDao      dao.User
	roleDao      dao.Role
	animeDao     dao.Anime
	categoryDao  dao.Category
	commentDao   dao.Comment
	messageDao   dao.Message
	biliInfoDao  dao.BiliInfo
	videoDao     dao.Video
	archiveDao   dao.Archive
	partitionDao dao.Partition
	adminDao     dao.Admin
)

var (
	BiliInfoService BiliInfo
)
