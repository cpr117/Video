// @User CPR
package utils

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"fmt"
)

func InitOther() {
	// 初始化B站的相关全局参数
	// 默认将其保存在数据库中
	biliConfigDetail := model.InitBiliConfig()
	fmt.Println(Json.Marshal(biliConfigDetail))
	biliConfig := model.BiliConfig{
		Config: Json.Marshal(biliConfigDetail),
	}
	dao.Create(&biliConfig)
}
