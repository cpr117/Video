// @User CPR
package model

type Partition struct {
	Universal
	Name string `gorm:"type:varchar(20);not null;unique" json:"name"`
}