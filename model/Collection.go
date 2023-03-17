// @User CPR
package model

import "reflect"

// 收藏夹
type Collection struct {
	Universal
	UserId    int    `gorm:"type:int(11);not null;comment:'用户ID'" json:"user_id"`
	Name      string `gorm:"type:varchar(255);not null;comment:'收藏夹名'" json:"name"`
	Count     int    `gorm:"type:int(11);not null;comment:'收藏夹视频数量'" json:"count"`
	IsPrivacy int8   `gorm:"type:tinyint(1);not null;comment:'是否私密(0-否 1-是)'" json:"is_privacy"`
}

func (c *Collection) IsEmpty() bool {
	return reflect.DeepEqual(c, &Collection{})
}
