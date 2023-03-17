// @User CPR
package dao

import (
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Role struct{}

// 根据[userInfoId] 获取角色标签列表
func (*Role) GetLabelsByUserInfoId(userInfoId int) (labels []string) {
	DB.Table("role r, user_role ur").
		Where("r.id = ur.role_id AND ur.user_id = ?", userInfoId).
		Pluck("label", &labels) // 将单列查询为切片
	return
}

// 根据[角色id] 获取角色标签列表
func (*Role) GetLabelsByRoleIds(ids []int) (labels []string) {
	DB.Model(model.Role{}).Where("id in ?", ids).Pluck("label", &labels)
	return
}

// 获取角色列表[非树形]
func (*Role) GetList(req req.PageQuery) (list []resp.RoleVo, total int64) {
	list = make([]resp.RoleVo, 0)
	db := DB.Table("role")
	// 查询条件
	if req.KeyWord != "" {
		db = db.Where("name like ?", "%"+req.KeyWord+"%")
	}
	db.Count(&total).
		Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).
		Select("id, name, label, created_at, is_disable").
		Find(&list)
	return list, total
}

// 获取角色选项
func (*Role) GetOption() []resp.OptionVo {
	list := make([]resp.OptionVo, 0)
	DB.Model(&model.Role{}).
		//Where("IsDisable", 0).
		Select("id", "name").Find(&list)
	return list
}

// 根据[ 角色id ] 获取 [资源id列表]
func (*Role) GetResourcesByRoleId(roleId int) (resourceIds []int) {
	DB.Model(&model.RoleResource{}).
		Where("role_id = ?", roleId).
		Pluck("resource_id", &resourceIds)
	return
}

//// 根据 [角色id] 查询出 [菜单id列表]
//func (*Role) GetMenusByRoleId(roleId int) (menuIds[] int){
//	DB.Model(&model.RoleMenu{}).
//		Where("role_id = ?", roleId).
//		Pluck("menu_id", &menuIds)
//	return
//}
