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
func BackRouter() http.Handler {
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
	base := r.Group("/v1")
	{
		// TODO: 后台登录
		// 这边默认最高用户在一开始就创建好了 下面的用户由高级后台用户向下创建 因此没有注册界面
		base.POST("/login", userAuthAPI.Login)
		base.POST("/register", userAuthAPI.Register)

		// 需要鉴权的借口
		auth := base.Group("/api")
		// !注意使用中间件的顺序
		auth.Use(middlewares.JWTAuth())      // JWT 鉴权中间件
		auth.Use(middlewares.RBAC())         // casbin 权限中间件
		auth.Use(middlewares.ListenOnline()) // 监听在线用户
		auth.Use(middlewares.OperationLog()) // 记录操作日志
		{
			//auth.GET("/home", biliInfoAPI.GetHomeInfo) // 后台首页信息
			auth.GET("/logout", userAuthAPI.Logout)      // 退出登录
			auth.POST("/register", userAuthAPI.Register) // 注册用户

			video := auth.Group("/video")
			{
				video.GET("/list", videoAPI.GetVideoList)
				//video.GET("/detail", videoAPI.DetailVideo)
				video.POST("/update", videoAPI.UpdateVideo)
				video.DELETE("/delete", videoAPI.DeleteVideo)
			}
			user := auth.Group("/user")
			{
				user.GET("/list", userAPI.GetList)
				//user.GET("/detail", userAPI.Detail)
				user.PUT("/disable/:id", userAPI.UpdateDisable)
				user.DELETE("/delete/:id", userAPI.DeleteUser)
				user.PUT("/role/:id", userAPI.UpdateRole)
			}
			admin := auth.Group("/admin")
			{
				admin.GET("/list", adminAPI.GetList)
				//admin.GET("/detail", adminAPI.Detail)
				admin.POST("/create", adminAPI.Create)
				admin.PUT("/update", adminAPI.Update)
				//admin.PUT("/disable", adminAPI.UpdateDisable)		// admin中暂时还是没有disable部分
				//admin.DELETE("/delete/:id", adminAPI.DeleteAdmin)
			}
			partition := auth.Group("/partition")
			{
				partition.POST("/add", partitionAPI.AddPartition)     // 添加新分区
				partition.GET("/list", partitionAPI.GetPartitionList) // 获取分区列表
			}
			comment := auth.Group("/comment")
			{
				comment.GET("/list/:videoid", commentAPI.GetListByVideoId) // 获取评论列表
				comment.DELETE("/delete/:id", commentAPI.DeleteComment)    // 删除评论

			}
		}
	}

	return r
}
