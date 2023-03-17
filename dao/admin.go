// @User CPR
package dao

import (
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Admin struct{}

func (*Admin) GetList(req req.GetAdminList) (ret []resp.AdminList, total int64) {
	db := DB.Model(&model.Admin{})
	if req.NickName != "" {
		db.Where("nick_name like ?", "%"+req.NickName+"%")
	}
	if req.Level != 0 {
		db.Where("level = ?", req.Level)
	}
	db.Select("id, nick_name, username, email, phone, role, level, create_time, update_time").
		Count(&total).
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Find(&ret)
	return
}
