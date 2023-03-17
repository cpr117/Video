// @User CPR
package model

type Admin struct {
	Universal
	UserName string `gorm:"column:username;type:varchar(20);not null;unique" json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Password string `gorm:"column:password;type:varchar(50);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string `gorm:"column:email;type:varchar(50);not null;unique" json:"email" validate:"required,email" label:"邮箱"`
	Phone    string `gorm:"column:phone;type:varchar(11);not null;unique" json:"phone" validate:"required,min=11,max=11" label:"手机号"`
	Role     int    `gorm:"column:role;type:int;not null" json:"role" validate:"required" label:"角色"`
	Level    int8   `gorm:"column:level;type:tinyint(1);not null" json:"level" validate:"required" label:"等级"`
}
