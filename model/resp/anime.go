// @User CPR
package resp

// 动漫列表Vo
type AnimeVo struct {
	Id            int      `json:"id"`             // 动漫ID
	CategoryNames []string `json:"category_names"` // 分类名

	Year       int     `json:"year"`        // 年份
	Season     int     `json:"season"`      // 季度
	Title      string  `json:"title"`       // 番名
	Desc       string  `json:"desc"`        // 简介
	Url        string  `json:"url"`         // 番剧链接
	Img        string  `json:"img"`         // 封面
	IsDelete   int8    `json:"is_delete"`   // 是否删除
	NeedMember int8    `json:"need_member"` // 是否需要会员
	Status     int8    `json:"status"`      // 状态 1:连载中 2:已完结
	UpdateTime string  `json:"update_time"` // 更新时间
	Score      float64 `json:"score"`       // 评分

	//WantCount int `json:"want_count"`	// 想看人数
	//ViewCount int `json:"view_count"`	// 看过人数

}

type AnimeRank struct {
	Id    int    `json:"id"`
	Score int    `json:"score"`
	Title string `json:"title"`
	Rank  int    `json:"rank"`
}
