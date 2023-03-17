// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
)

type BiliInfo struct{}

// Cache Aside Pattern: 经典的缓存 + 数据库的读写模式

// 获取设置
func (*BiliInfo) GetBiliConfig() (respVo model.BiliConfigDetail) {
	// 尝试从Redis中取值
	biliConfig := utils.Redis.GetVal(KEY_BILI_CONFIG)
	// Redis 中没有值， 再查数据库 查到后设置到Redis中
	if biliConfig == "" {
		biliConfig = dao.GetOne[model.BiliConfig]("id", 1).Config
		utils.Redis.Set(KEY_BILI_CONFIG, biliConfig, 0)
	}
	// 反序列化字符串为golang对象
	utils.Json.Unmarshal(biliConfig, &respVo)
	return respVo
}

// 获取后端首页信息
//func (*BiliInfo) GetBackHomeInfo() resp.BiliBackHomeVo {
//
//}

// 获取前端首页信息
func (*BiliInfo) GetFrontHomeInfo() (ret resp.BiliFrontHomeVo) {
	list := biliInfoDao.GetFrontHomeInfo()
	for i := 0; i < len(list); i++ {
		list[i].Rank = i + 1
	}
	ret.MonthHotRank = list
	return
}
