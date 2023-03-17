// @User CPR
package model

// 用户信息只包含了用户 个人主页内显示的信息
type User struct {
	Universal
	Coins    int  `gorm:"int(11);comment:硬币数" json:"coins"`
	IsMember bool `gorm:"type:bool;comment:是否是会员" json:"is_member"`

	//Followers    []int  `gorm:"type:json;comment:粉丝id列表" json:"followers"`
	//FavorVideos  []int  `gorm:"type:json;comment:收藏视频id列表" json:"favor_videos"`
	//FansIds      []int  `gorm:"type:json;comment:关注的用户id列表" json:"fans_ids"`
	Avatar       string `gorm:"type:varchar(255);not null;comment:头像" json:"avatar"`
	NickName     string `gorm:"type:varchar(255);not null;comment:昵称" json:"nick_name"`
	PersonalSign string `gorm:"type:varchar(255);not null;comment:个性签名" json:"personal_sign"`
	//SubList      []Sub     `gorm:"type:varchar(255);not null;comment:上传列表" json:"sub_list"`
	IsDisable int8 `gorm:"type:tinyint(1);comment:是否禁用(0-否 1-是)" json:"is_disable"`
}

//type UserAuth struct{
//	Universal
//	UserName string `gorm:"type:varchar(255);not null;comment:用户名" json:"user_name"`
//	Password string `gorm:"type:varchar(255);not null;comment:密码" json:"password"`
//	Email   string `gorm:"type:varchar(255);not null;comment:邮箱" json:"email"`
//}

// --------关联表-----------

type UserFollow struct {
	UserId   int `gorm:"int(11);comment:用户id" json:"user_id"`
	FollowId int `gorm:"int(11);comment:关注的用户id" json:"follow_id"`
}

type UserFans struct {
	UserId int `gorm:"int(11);comment:用户id" json:"user_id"`
	FanId  int `gorm:"int(11);comment:粉丝id" json:"fans_id"`
}

type UserFavor struct {
	UserId  int `gorm:"int(11);comment:用户id" json:"user_id"`
	VideoId int `gorm:"int(11);comment:视频id" json:"video_id"`
}

type UserCollection struct {
	UserId       int `gorm:"int(11);comment:用户id" json:"user_id"`
	CollectionId int `gorm:"int(11);comment:收藏夹id" json:"collection_id"`
}
