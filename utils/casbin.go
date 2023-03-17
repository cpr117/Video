// @User CPR
package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"sync"
)

const CASBIN_UTIL_ERR_PREFIX = "utils/casbin.go -> "

var (
	cachedEnforcer *casbin.CachedEnforcer
	casbin_db      *gorm.DB
	once           sync.Once
)

type _casbin struct{}

var Casbin = new(_casbin)

func InitCasbin(db *gorm.DB) *casbin.CachedEnforcer {
	var err error
	once.Do(func() {
		casbin_db = db

		adapter, _ := gormadapter.NewAdapterByDB(db)

		// 方法一: 从字符串中加载
		text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
		m, _ := model.NewModelFromString(text)
		cachedEnforcer, err = casbin.NewCachedEnforcer(m, adapter)
		if err != nil {
			log.Panic("Casbin 初始化失败: ", err)
		}

		// 方法二: 从配置文件中加载
		// cachedEnforcer, _ = casbin.NewCachedEnforcer("config/casbin.conf", adapter)

		//InitRoleData()

		cachedEnforcer.SetExpireTime(60 * 60)
		cachedEnforcer.EnableAutoSave(true)
		cachedEnforcer.LoadPolicy()

		log.Println("Casbin 初始化成功")
	})

	return cachedEnforcer
}

func InitRoleData() error {
	ok, err := cachedEnforcer.AddPolicies([][]string{
		{"admin", "/v1/api/video/list", "GET"},
		{"admin", "/v1/api/video/update", "POST"},
		{"admin", "/v1/api/video/delete", "DELETE"},

		{"admin", "/v1/api/user/list", "GET"},
		{"admin", "/v1/api/user/disable", "PUT"},
		{"admin", "/v1/api/user/delete", "DELETE"},
		{"admin", "/v1/api/user/role", "PUT"},

		{"admin", "/v1/api/admin/list", "GET"},
		{"admin", "/v1/api/admin/create", "POST"},
		{"admin", "/v1/api/admin/update", "PUT"},

		{"admin", "/v1/api/partition/add", "POST"},
		{"admin", "/v1/api/partition/list", "GET"},

		{"admin", "/v1/api/comment/list", "GET"},
	})
	if ok == false {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"InitRoleData: ", zap.Error(err))
		return err
	}
	return nil
}

func (*_casbin) Enforcer() *casbin.CachedEnforcer {
	return cachedEnforcer
}

// RBAC API
func (*_casbin) AddRoleForUser(user, role string) {
	_, err := cachedEnforcer.AddRoleForUser(user, role)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"AddRoleForUser: ", zap.Error(err))
		panic(err)
	}
}

// 删除用户角色
// func (c *Casbin) DeleteRoleForUser(user, role string) {
// 	_, err := cachedEnforcer.DeleteRoleForUser(user, role)
// 	if err != nil {
// 		log.Println(CASBIN_UTIL_ERR_PREFIX+"DeleteRoleForUser: ", err)
// 		panic(err)
// 	}
// }

// 添加 API 权限: AddPermissionForRole("user", "/article/list", "GET")
func (c *_casbin) AddPermissionForRole(role string, params ...string) {
	_, err := cachedEnforcer.AddPermissionForUser(role, params...)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"AddPermissionForRole: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 批量添加 Policy
func (c *_casbin) AddPolicies(rules [][]string) {
	if len(rules) == 0 {
		return
	}
	_, err := cachedEnforcer.AddPolicies(rules)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"AddPolicies: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 删除API权限
func (c *_casbin) DeletePermission(permission ...string) {
	_, err := cachedEnforcer.DeletePermission(permission...)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"DeletePermissison: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 删除角色的API权限
func (c *_casbin) DeletePermissionForRole(role string, permissions ...string) {
	_, err := cachedEnforcer.DeletePermissionForUser(role, permissions...)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"DeletePermissionForRole: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 批量删除 API 权限: 根据角色名
func (c *_casbin) BatchDeletePermissions(roles []string) {
	for _, role := range roles {
		_, err := cachedEnforcer.DeletePermissionsForUser(role) // 删除改角色对应的权限 API
		if err != nil {
			Logger.Error(CASBIN_UTIL_ERR_PREFIX+"BatchDeletePermissions: ", zap.Error(err))
			panic(err)
		}
	}
	// err := casbin_db.Exec("DELETE FROM casbin_rule WHERE v0 in ?", labels).Error
	// if err != nil {
	// 	log.Println(CASBIN_UTIL_ERR_PREFIX+"RemovePoliciesByLabels: ", err)
	// 	panic(err)
	// }
	c.LoadPolicy()
}

// 更新 API 权限
func (c *_casbin) UpdatePolicy(old []string, new []string) {
	_, err := cachedEnforcer.UpdatePolicy(old, new)
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"UpdatePolicy: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 更新 API 权限: 修改角色名
func (c *_casbin) UpdateRoleLabels(oldLabel, newLabel string) {
	err := casbin_db.Model(&gormadapter.CasbinRule{}).
		Where("v0 = ?", oldLabel).Update("v0", newLabel).Error
	if err != nil {
		Logger.Error(CASBIN_UTIL_ERR_PREFIX+"UpdateRoleLabels: ", zap.Error(err))
		panic(err)
	}
	c.LoadPolicy()
}

// 重新加载策略, 使得更新的策略生效
func (*_casbin) LoadPolicy() {
	// log.Println("重新加载策略, 使得更新的策略生效...")
	cachedEnforcer.LoadPolicy()
}
