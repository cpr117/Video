// @User CPR
package resp

import "time"

// 角色 + 资源id列表 + 菜单id列表
type RoleVo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`                  // 角色名
	Label       string    `json:"label"`                 // 角色标签
	ResourceIds []int     `json:"resource_ids" gorm:"-"` // 资源id列表
	MenuIds     []int     `json:"menu_ids" gorm:"-"`     // 菜单id列表
	IsDisable   int       `json:"is_disable"`
	CreatedAt   time.Time `json:"created_at"`
}

// 后台
type GetUser struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      int       `json:"role"`
	IsDisable int       `json:"is_disable"`
	Avatar   string    `json:"avatar"`

}
