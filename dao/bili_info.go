// @User CPR
package dao

import (
	"VideoWeb/model/resp"
)

type BiliInfo struct{}

func (*BiliInfo) GetFrontHomeInfo() (list []resp.AnimeRank) {
	DB.Table("anime a").
		Select("a.id, a.title, a.score").
		Order("a.score DESC").
		Limit(100).
		Find(&list)
	return
}
