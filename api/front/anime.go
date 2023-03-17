// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Anime struct{}

func (*Anime) GetAnimeList(c *gin.Context) {
	r.SuccessData(c, animeService.GetList(utils.BindQuery[req.GetAnimes](c)))
}

func (*Anime) GetAnime(c *gin.Context) {
	r.SuccessData(c, animeService.GetAnime(utils.BindJson[req.GetAnime](c).AnimeId))
}

func (*Anime) GetAnimeRank(c *gin.Context) {
	r.SuccessData(c, animeService.GetAnimeRank(utils.BindQuery[req.GetAnimeRank](c)))
}

// 点赞动漫
func (*Anime) LikeAnime(c *gin.Context) {
	uid := utils.GetFromContext[int](c, "user_info_id")
	animeId := utils.GetIntParam(c, "anime_id")
	r.SendCode(c, animeService.LikeAnime(uid, animeId))
}
