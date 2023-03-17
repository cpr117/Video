// @User CPR
package model

import "reflect"

type Category struct {
	Universal
	Name string `gorm:"type:varchar(100);not null;comment:分类名称" json:"name"`
	//Animes []Anime `gorm:"foreignKey:CategoryId"` // 重写外键
	Count int `gorm:"-" json:"count"`
}

func (c *Category) IsEmpty() bool {
	return reflect.DeepEqual(c, &Category{})
}
