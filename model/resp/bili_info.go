// @User CPR
package resp

import "MyBilibili/model"

// 后台首页Vo
type BiliBackHomeVo struct {
	VideoCount    int64                  `json:"video_count"`    // 视频总数
	UserCount     int64                  `json:"user_count"`     // 用户总数
	ViewCount     int64                  `json:"view_count"`     // 访问量
	CategoryCount int64                  `json:"category_count"` // 分类总数
	BlogConfig    model.BiliConfigDetail `json:"bili_config"`    // B站信息
}

// 前台首页Vo
type BiliFrontHomeVo struct {
	DayHotRank   []AnimeRank `json:"day_hot_rank"`   // 日热榜
	WeekHotRank  []AnimeRank `json:"week_hot_rank"`  // 周热榜
	MonthHotRank []AnimeRank `json:"month_hot_rank"` // 月热榜
}
