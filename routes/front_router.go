// @User CPR
package routes

import (
	"VideoWeb/config"
	"VideoWeb/routes/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 后台管理界面的接口路由
func FrontRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	// 使用本地文件上传 需要静态文件服务
	if config.Cfg.Upload.OssType == "local" {
		r.Static("/public", "./public")
		r.StaticFS("/dir", http.Dir("./public")) // 将 public 目录内的文件列举展示
	}

	r.Use(middlewares.Logger())             // 自定义的zap日志中间件
	r.Use(middlewares.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middlewares.Cors())               // 跨域中间件

	// 基于 cookie 存储 session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt))

	// session 存储时间和JWT（json web token）过期时间一致
	// 这是两种不同的用户检测机制 一者存储在客户端 一者存储在服务端
	store.Options(sessions.Options{MaxAge: int(config.Cfg.JWT.Expire) * 3600})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store)) // 使用session中间件

	// 无需鉴权的借口
	base := r.Group("")
	{
		// TODO: 首页 用户登录 用户注册 直播主页 直播间 视频界面 视频分类
		base.POST("/login", fUserAuthAPI.Login)
		base.POST("/register", fUserAuthAPI.Register) // 注册 (包含邮箱验证码实现)
		base.POST("/code", fUserAuthAPI.GetCaptcha)   // 获取验证码
		//base.GET("/home", fBiliInfoAPI.GetHomeInfo)

		base.GET("/anime_cate", fCateAPI.GetCategoryList) // 获取分类信息
		//base.GET("/rand_video", RandVedio)			// 获取随机视频

		// 动漫
		anime := r.Group("/anime")
		{
			anime.GET("/animeList", fAnimeAPI.GetAnimeList) // 获取视频信息
			anime.GET("/:id", fAnimeAPI.GetAnime)
			anime.GET("/rank", fAnimeAPI.GetAnimeRank)
		}
		// video
		video := r.Group("/video")
		{
			video.POST("/upload", fVideoAPI.UploadVideo)
			//video.GET("/list", fVideoAPI.GetVideoList)
			video.GET("/:id", fVideoAPI.GetVideoById) // video
			video.DELETE("/:id", fVideoAPI.DeleteVideo)
		}
		partition := r.Group("/partition")
		{
			partition.GET("/list", fPartitionAPI.GetPartitionList)
		}
		// TODO:直播
		//live := r.Group("/live")
		//{
		//	live.GET("/liveList", liveAPI.GetLiveList)
		//	live.GET("/cate", liveAPI.LiveCate)
		//	live.GET("/room/:id", LiveRoom)
		//}
		comment := r.Group("/comment")
		{
			comment.GET("/list", fCommentAPI.GetFrontList)
		}
	}

	// 需要鉴权的借口
	base.Use(middlewares.JWTAuth())
	auth := base.Group("/auth")
	{
		user := auth.Group("/user")
		{
			user.GET("/logout", fUserAuthAPI.Logout) // 退出登录
			user.GET("/info", fUserAPI.GetInfo)      // 获取用户信息
			//user.POST("update_pass", fUserAPI.UpdatePassWord) // 修改密码

		}

		comment := auth.Group("/comment")
		{
			//comment.GET("/list", fCommentAPI.GetFrontList)
			comment.POST("/reply/:comment_id", fCommentAPI.ReplyComment)
			comment.DELETE("/reply/:comment_id", fCommentAPI.DeleteComment)

		}
		message := auth.Group("/message")
		{
			message.GET("/list", fMessageAPI.GetMessageList)
			message.POST("/send", fMessageAPI.SendMessage)
		}

		anime := auth.Group("/anime")
		{
			anime.POST("/like/:id", fAnimeAPI.LikeAnime)
		}
		archive := auth.Group("/collection")
		{
			archive.GET("/collection/list", fCollectionAPI.GetCollectionList)
			archive.POST("/collection/add", fCollectionAPI.AddToCollection)
			archive.POST("/collection/del", fCollectionAPI.DelFromCollection)
			archive.POST("/collection/create", fCollectionAPI.CreateCollection)
			archive.POST("/collection/update", fCollectionAPI.UpdateCollection)
			archive.DELETE("/collection/delete/:id", fCollectionAPI.DeleteCollection)
		}
	}
	return r
}
