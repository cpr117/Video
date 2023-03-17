// @User CPR
package model

import "reflect"

const (
	PUBLIC  = iota + 1 // 公开
	BLOCK              // 屏蔽
	EXAMINE            // 审核
)

//1、通过HTML<video>标签可以播放指定的视频
//2、前后端分离,接口只返回视频播放地址
//3、使用Golang语言的gin框架实现
//实现思路：
//1、Gin框架做接口端,返回视频文件（本质上就是下载视频文件接口）
//2、<video>标签中的src地址填写视频下载地址,通过Chrome在线播放

type Anime struct {
	Universal
	Title       string  `gorm:"type:varchar(255);not null;comment:'动漫标题'" json:"title"`
	Img         string  `gorm:"type:varchar(255);not null;comment:'动漫封面'" json:"img"`
	CategoryIds []int   `gorm:"type:bigint;not null;comment:分类 ID" json:"category_id"`
	Desc        string  `gorm:"type:varchar(200);comment:描述" json:"desc"`
	Avatar      string  `gorm:"type:varchar(100);comment:封面图片地址" json:"avatar"`
	Year        int     `gorm:"type:int;comment:年份" json:"year"`
	Season      int     `gorm:"type:int;comment:季度" json:"season"`
	Episode     int     `gorm:"type:int;comment:集数" json:"episode"`
	IsDelete    int8    `gorm:"type:tinyint;default:0;comment:是否删除" json:"is_delete"`
	Score       float64 `gorm:"type:decimal(3,1);comment:评分" json:"score"`
	UpdateTime  string  `gorm:"type:datetime;comment:更新时间" json:"update_time"`
	NeedMember  int8    `gorm:"type:tinyint;default:0;comment:是否需要会员" json:"need_member"`
	Url         string  `gorm:"type:varchar(255);comment:番剧链接" json:"url"`
}

type AnimeResource struct {
	AnimeId  int    `gorm:"type:int(11);not null;comment:视频id" json:"video_id"`
	Res360   string `gorm:"type:varchar(255);not null;comment:360P视频地址" json:"res_360"`
	Res480   string `gorm:"type:varchar(255);not null;comment:480P视频地址" json:"res_480"`
	Res720   string `gorm:"type:varchar(255);not null;comment:720P视频地址" json:"res_720"`
	Res1080  string `gorm:"type:varchar(255);not null;comment:1080P视频地址" json:"res_1080"`
	Duration string `gorm:"type:varchar(255);not null;comment:视频时长" json:"duration"`
}

func (a *Anime) IsEmpty() bool {
	return reflect.DeepEqual(a, &Anime{})
}

// 动漫 - 分类 关联
type AnimeCategory struct {
	AnimeId    int
	CategoryId int
}
