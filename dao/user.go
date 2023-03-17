// @User CPR
package dao

import (
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type User struct{}

func (*User) GetFollows(id int) (follows []int) {
	DB.Table("UserFollow uf").
		Select("uf.follow_id").
		Where("uf.user_id = ?", id).
		Find(&follows)
	return
}

func (*User) GetFans(id int) (fans []int) {
	DB.Table("UserFan uf").
		Select("uf.fan_id").
		Where("uf.user_id = ?", id).
		Find(&fans)
	return
}

func (*User) GetFavor(id int) (favors []int) {
	DB.Table("UserFavor uf").
		Select("uf.favor_id").
		Where("uf.user_id = ?", id).
		Find(&favors)
	return
}

// 后台
func (*User) GetUserList(req req.GetUserList) (data []resp.UserDetailVo, total int64) {
	//model.User{}
	//model.UserAuth{}
	db := DB.Table("User u")
	if req.NickName != "" {
		db = db.Where("u.nick_name like ?", "%"+req.NickName+"%")
	}
	db.Joins("left join user_auth ua ON u.id = ua.user_info_id").
		Select("u.id, u.nick_name, u.avatar, u.personal_sign, u.coins, u.is_member, u.is_disable, ua.email, ua.ip_address, ua.ip_source, ua.last_login_time").
		Find(&data)
	return
}
