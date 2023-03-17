// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"strconv"
	//"VideoWeb/dao"
)

type Anime struct{}

func (*Anime) GetList(req req.GetAnimes) (ret resp.PageResult[[]resp.AnimeVo]) {
	animeList, total := animeDao.GetList(req)

	//wantCountMap := utils.Redis.HGetAll(KEY_ANIME_WANT_COUNT)	// 想看人数 map
	//viewCountZ := utils.Redis.ZRangeWithScores(KEY_ANIME_VIEW_COUNT, 0, -1)
	//viewCountMap := getViewCountMap(viewCountZ)	// 访问数量 map

	//for i, anime := range animeList {
	//	animeList[i].ViewCount = viewCountMap[anime.ID]
	//	animeList[i].WantCount, _ = strconv.Atoi(wantCountMap[strconv.Itoa(anime.ID)])
	//}

	ret.PageSize = req.PageSize
	ret.PageNum = req.PageNum
	ret.Total = total
	ret.List = animeList
	return
}

// 动漫检索
func (*Anime) GetAnime(id int) resp.AnimeVo {
	anime := dao.GetOne[model.Anime]("id", id)
	categories := categoryDao.GetCategoryNamesByAnimeId(id)

	animeVo := utils.CopyProperties[resp.AnimeVo](anime)
	// 前端 category 为 '' 不显示 placeholder, 为 null 显示 placeholder
	animeVo.CategoryNames = categories
	return animeVo
}

func (*Anime) GetAnimeRank(req req.GetAnimeRank) (ret resp.PageResult[[]resp.AnimeVo]) {
	animeList, total := animeDao.GetAnimeRank(req)

	ret.PageSize = req.PageSize
	ret.PageNum = req.PageNum
	ret.Total = total
	ret.List = animeList
	return ret
}

func (*Anime) LikeAnime(userId int, animeId int) (code int) {
	// 用户已经点过赞
	animeLikeUserKey := KEY_ANIME_USER_LIKE_SET + strconv.Itoa(userId)
	if utils.Redis.SIsMember(animeLikeUserKey, animeId) {
		utils.Redis.SRem(animeLikeUserKey, animeId)
		utils.Redis.HIncrBy(KEY_ANIME_LIKE_COUNT, strconv.Itoa(animeId), -1)
	} else {
		utils.Redis.SAdd(animeLikeUserKey, animeId)
		utils.Redis.HIncrBy(KEY_ANIME_LIKE_COUNT, strconv.Itoa(animeId), 1)
	}
	return r.OK
}
