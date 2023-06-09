// @User CPR
package model

import "time"

type Universal struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}

type PageQuery struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Keyword  string `form:"keyword"`
}
