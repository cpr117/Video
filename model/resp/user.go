// @User CPR
package resp

import "time"

// LoginVo
// 登录返回信息 一些比较简洁的信息
// 例如私信 并不是获取全部信息 而是 获取未读信息有多少
type LoginVo struct {
	Id         int `json:"id"`
	UserInfoId int `json:"user_info_id"` // 用户信息ID

	Email        string `json:"email"` // 邮箱
	NickName     string `json:"nickname"`
	Coins        int    `json:"coins"`         // 硬币数量
	IsMember     bool   `json:"is_member"`     // 是否是大会员
	Avatar       string `json:"avatar"`        // 头像
	PersonalSign string `json:"personal_sign"` // 个性签名

	FollowNum int `json:"follow_num"` // 关注数量
	FansNum   int `json:"fans_num"`   // 粉丝数量

	SubNum    int `json:"sub_num"`    // 投稿数量 数据库count 一下
	UnreadMsg int `json:"unread_msg"` // 未读消息数量 数据库count 一下

	// todo: 动态相关的内容
	// UnreadEvent int `json:"unread_event"` // 未读动态数量 数据库count 一下
	// EventNum    int `json:"event_num"`    // 自己发的动态数量 数据库count 一下

	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	LastLoginTime time.Time `json:"last_login_time"`
	LoginType     int       `json:"login_type"`

	Token string `json:"token"`
}

type UserDetailVo struct {
	UserId       int    `json:"user_info_id"` // 用户信息ID
	Coins        int    `json:"coins"`        // 硬币数量
	NickName     string `json:"nickname"`
	IsMember     bool   `json:"is_member"`     // 是否是大会员
	Avatar       string `json:"avatar"`        // 头像
	Email        string `json:"email"`         // 邮箱
	PersonalSign string `json:"personal_sign"` // 个性签名
	FansList     []int  `json:"fans_list"`     // 粉丝列表
	FollowList   []int  `json:"follow_list"`   // 关注列表
	FavorList    []int  `json:"favor_video"`   // 收藏的视频
}

type UserVo struct {
	UserId   int    `json:"user_info_id"` // 用户信息ID
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"` // 头像
}

// 后台
type BUserDetailVo struct {
	UserId        int       `json:"user_info_id"` // 用户信息ID
	Coins         int       `json:"coins"`        // 硬币数量
	NickName      string    `json:"nickname"`
	UserName      string    `json:"user_name"`
	IsMember      bool      `json:"is_member"`     // 是否是大会员
	Avatar        string    `json:"avatar"`        // 头像
	Email         string    `json:"email"`         // 邮箱
	PersonalSign  string    `json:"personal_sign"` // 个性签名
	IpAddress     string    ` json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	LastLoginTime time.Time `json:"last_login_time"`
}
