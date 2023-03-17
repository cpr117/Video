// @User CPR
package dao

import (
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Anime struct{}

// 前台接口
func (*Anime) GetList(req req.GetAnimes) ([]resp.AnimeVo, int64) {
	var list = make([]resp.AnimeVo, 0)
	var total int64

	db := DB.Model(model.Anime{})
	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}
	//if req.CategoryId != 0 {
	//	db = db.Where("? in category_id", req.CategoryId)
	//}
	if req.Year != 0 {
		db = db.Where("year", req.Year)
	}
	if req.Status != 0 {
		db = db.Where("status", req.Status)
	}
	db.Preload("Category").
		Joins("LEFT JOIN anime_category ON anime.id = anime_category.anime_id AND "+
				"anime_category.category_id = ?", req.CategoryId).
		Group("id"). // 去重
		Count(&total).
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Order("is_top DESC, anime.score DESC, anime.id DESC").
		Find(&list)
	return list, total
}

func (*Anime) GetAnimeById(req req.GetAnime) resp.AnimeVo {
	var anime resp.AnimeVo
	DB.Where("id", req.AnimeId).
		First(&anime)
	return anime
}

func (*Anime) GetAnimeRank(req req.GetAnimeRank) (ret []resp.AnimeVo, total int64) {
	db := DB.Model(model.Anime{})
	if req.KeyWord != "" {
		db = db.Where("title LIKE ? OR desc LIKE ?", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%")
	}
	if req.Year != 0 {
		db = db.Where("year", req.Year)
	}
	if req.Season != 0 {
		db = db.Where("season", req.Season)
	}
	if req.CategoryId != 0 {
		db = db.Where("id IN (SELECT anime_id FROM anime_category WHERE category_id = ?)", req.CategoryId)
	}
	db.Count(&total).
		Order("score DESC, id DESC").
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Find(&ret)
	return
}
