// @User CPR
package model

// JSON格式的B站配置
type BiliConfig struct {
	Universal
	Config string `gorm:"type:varchar(2000);not null;comment:配置" json:"config"`
}

type BiliConfigDetail struct {
	// 网站信息
	WebsiteAvatar     string `gorm:"" json:"website_avatar" label:"网站头像"`
	WebsiteName       string `json:"website_name" label:"网站名称"`
	WebsiteIntro      string `json:"website_intro" label:"网站简介"`
	WebsiteAuthor     string `json:"website_author" label:"网站作者"`
	WebsiteNotice     string `json:"website_notice" label:"网站公告"`
	WebsiteCreateTime string `json:"website_create_time" label:"网站创建时间"`
	WebsiteRecordNo   string `json:"website_record" label:"网站备案号"`

	// 社交信息
	SocialLoginList []string `json:"social_login_list" label:"第三方登录"`
	SocialUrlList   []string `json:"social_url_list" label:"登录地址"`
	Qq              string   `json:"qq" label:"QQ"`
	Github          string   `json:"github" label:"Github"`
	Gitee           string   `json:"gitee" label:"Gitee"`

	// 其他设置
	TouristAvatar   string `json:"tourist_avatar"`    // 默认游客头像
	UserAvatar      string `json:"user_avatar"`       // 默认用户头像
	ArticleCover    string `json:"article_cover"`     // 默认文章封面
	IsCommentReview int8   `json:"is_comment_review"` // 评论默认审核
	IsMessageReview int8   `json:"is_message_review"` // 留言默认审核
	IsEmailNotice   int8   `json:"is_email_notice"`   // 邮箱通知
	IsReward        int8   `json:"is_reward"`         // 打赏状态
	WeiXinQRCode    string `json:"wechat_qrcode"`     // 微信收款码图片
	AlipayQRCode    string `json:"alipay_ode"`        // 支付宝收款码图片
	// IsChatRoom      int    `json:"is_chatroom"`       // 聊天室状态
	// WebsocketUrl    string `json:"websocket_url"`     // websocket 地址
	// IsMusicPlayer   int    `json:"is_music_player"`   // 音乐播放器状态
}

func InitBiliConfig() BiliConfigDetail {
	biliConfigDetail := BiliConfigDetail{
		// 网站信息
		WebsiteAvatar:     "",
		WebsiteName:       "VideoWeb",
		WebsiteIntro:      "这是一个视频网站 试图模仿B站",
		WebsiteAuthor:     "cpr",
		WebsiteNotice:     "",
		WebsiteCreateTime: "2023-02-21",
		WebsiteRecordNo:   "京ICP备00000000号",

		// 社交信息
		SocialLoginList: []string{"QQ", "Github", "Gitee"},
		SocialUrlList:   []string{"https://www.qq.com", "https://www.github.com", "https://www.gitee.com"},

		// 其他设置
		TouristAvatar:   "https://www.baidu.com/img/bd_logo1.png",
		UserAvatar:      "https://www.baidu.com/img/bd_logo1.png",
		ArticleCover:    "https://www.baidu.com/img/bd_logo1.png",
		IsCommentReview: 0,
		IsMessageReview: 0,
		IsEmailNotice:   0,
		IsReward:        0,
		WeiXinQRCode:    "https://www.baidu.com/img/bd_logo1.png",
		AlipayQRCode:    "https://www.baidu.com/img/bd_logo1.png",
		// IsChatRoom:      0,
		// WebsocketUrl:    "",
		// IsMusicPlayer:   0,
	}

	return biliConfigDetail
}
