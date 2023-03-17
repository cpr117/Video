// @User CPR
package dao

import (
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Category struct{}

func (*Category) GetCategoryList(req req.PageQuery) (data []resp.CategoryVo, total int64) {
	db := DB.Table("category c").
		Select("c.id", "c.name", "COUNT(a.id) AS anime_count", "c.created_at", "c.updated_at").
		Joins("LEFT JOIN anime a ON c.id = a.category_id AND a.is_delete = 0 AND a.status = 1")
	if req.KeyWord != "" {
		db.Where("name LIKE ?", "%"+req.KeyWord+"%")
	}

	db.Count(&total).
		Order("id DESC").
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Find(&data)
	return
}

func (*Category) GetCategoryNamesByAnimeId(animeId int) (names []string) {

	DB.Table("category c").
		Select("c.name").
		Joins("LEFT JOIN anime_category ac ON c.id = ac.category_id").
		Where("ac.anime_id", animeId).
		Find(&names)
	return
}
