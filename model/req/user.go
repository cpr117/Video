// @User CPR
package req

// 前台
type Login struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	PassWord string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

type Logout struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
}

type Register struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Code     string `json:"code" validate:"required,min=6,max=6" label:"验证码"`
	Phone    string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
}

type RegisterCode struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
}

// 后台
type GetUserList struct {
	PageQuery
	NickName string `json:"nick_name" validate:"omitempty,min=2,max=20" label:"昵称"`
}

type UpdateDisable struct {
	ID        int  `json:"id" validate:"required" label:"id"`
	IsDisable int8 `json:"is_disable" validate:"required" label:"是否禁用"`
}

type V1Register struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Phone    string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
	Role     int8   `json:"role" validate:"required" label:"角色"`
}

type UpdateRole struct {
	ID      int `json:"id" validate:"required" label:"id"`
	AdminID int `json:"admin_id" validate:"required" label:"admin_id"`
	RoleID  int `json:"role_id" validate:"required" label:"role_id"`
}
