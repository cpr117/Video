// @User CPR
package req

type AddPartition struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePartition struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
