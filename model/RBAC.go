// @User CPR
package model

import (
	"reflect"
	"time"
)

// 权限控制
// 角色
type Role struct {
	Universal
	Name      string `gorm:"type:varchar(20);not null;comment:角色名" json:"name"`
	Label     string `gorm:"type:varchar(50);not null;comment:角色标签" json:"label"`
	IsDisable int    `gorm:"type:tinyint(1);not null;default:0;comment:是否禁用(0-否 1-是)" json:"is_disable"`
}

func (r *Role) IsEmpty() bool {
	return reflect.DeepEqual(r, &Role{})
}

// 资源(接口)
// Resource == casbin_rule
type Resource struct {
	Universal     `mapstructure:",squash"`
	Url           string `gorm:"type:varchar(255);comment:资源路径(接口URL)" json:"url" mapstructure:"url"`
	RequestMethod string `gorm:"type:varchar(10);comment:请求方式" json:"request_method" mapstructure:"request_method"`
	Name          string `gorm:"type:varchar(50);comment:资源名(接口名)" json:"name" mapstructure:"name"`
	ParentId      int    `gorm:"comment:父权限id" json:"parent_id" mapstructure:"parent_id,omitempty"`
	IsAnonymous   int    `gorm:"type:tinyint(1);comment:是否匿名访问(0-否 1-是)" json:"is_anonymous" mapstructure:"is_anonymous"`
}

// 角色-资源 关联
type RoleResource struct {
	RoleId     int `gorm:"type:int(11);primary_key" json:"role_id" mapstructure:"role_id"`
	ResourceId int `gorm:"type:int(11)" json:"resource_id" mapstructure:"resource_id"`
}

// 用户-角色 关联
type UserRole struct {
	UserId int `gorm:"primary_key" json:"user_id" mapstructure:"user_id"`
	RoleId int `gorm:"type:int(11)" json:"role_id" mapstructure:"role_id"`
}

//// 菜单
//type Menu struct {
//	Universal `mapstructure:",squash"`
//	Name      string `gorm:"type:varchar(20);comment:菜单名" json:"name"`
//	Path      string `gorm:"type:varchar(50);comment:菜单路径" json:"path"`
//	Component string `gorm:"type:varchar(50);comment:组件路径" json:"component"`
//	Icon      string `gorm:"type:varchar(50);comment:菜单图标" json:"icon"`
//	ParentId  int    `gorm:"comment:父菜单id" json:"parent_id"`
//	OrderNum  int8   `gorm:"type:tinyint(1);comment:排序" json:"order_num"`
//	IsHidden  int8   `gorm:"type:tinyint(1);comment:是否隐藏(0-否 1-是)" json:"is_hidden"`
//	KeepAlive int8   `gorm:"type:tinyint(1);default:1" json:"keep_alive"`
//	Redirect  string `gorm:"type:varchar(50);comment:重定向" json:"redirect"`
//}
//
//// 角色-菜单 关联
//type RoleMenu struct{
//	RoleId int `json:"role_id"`
//	MenuId int `json:"menu_id"`
//}

// 用户账户信息
type UserAuth struct {
	Universal     `mapstructure:",squash"`
	UserInfoId    int       `gorm:"comment:用户信息ID" json:"user_info_id"`
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(100);comment:密码" json:"password"`
	Email         string    `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	LoginType     int8      `gorm:"type:tinyint(1);comment:登录类型" json:"login_type"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间" json:"last_login_time"`
}
