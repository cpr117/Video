// @User CPR
package req

// 条件查询动漫
type GetAnimes struct {
	PageQuery
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
	Year       int    `form:"year"`
	//TagId int `form:"tag_id"`
	Status   int8  `form:"status" validate:"required,min=1,max=3"`
	IsDelete *int8 `form:"is_delete" validate:"required,min=0,max=1"`
}

// 条件查询动漫
type GetAnime struct {
	AnimeId int `json:"anime_id"`
}

// 获取分类动漫排行
type GetAnimeRank struct {
	PageQuery
	Year       int `json:"year"`
	Season     int `json:"season"`
	CategoryId int `json:"category_id"`
}
