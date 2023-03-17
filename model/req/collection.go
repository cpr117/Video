// @User CPR
package req

type GetCollection struct {
	PageQuery
	UserId       int `form:"user_id" validate:"required" json:"user_id"`
	CollectionId int `form:"collection_id" validate:"required" json:"collection_id"`
}

type AddToCollection struct {
	UserId       int `form:"user_id" validate:"required" json:"user_id"`
	CollectionId int `form:"collection_id" validate:"required" json:"collection_id"`
	VideoId      int `form:"video_id" validate:"required" json:"video_id"`
}

type CreateCollection struct {
	UserId         int    `form:"user_id" validate:"required" json:"user_id"`
	CollectionId   int    `form:"collection_id" validate:"required" json:"collection_id"`
	CollectionName string `form:"collection_name" validate:"required" json:"collection_name"`
	IsPrivacy      int8   `form:"is_privacy" validate:"required" json:"is_privacy"`
}

type DelFromCollection struct {
	UserId       int `form:"user_id" validate:"required" json:"user_id"`
	CollectionId int `form:"collection_id" validate:"required" json:"collection_id"`
	VideoId      int `form:"video_id" validate:"required" json:"video_id"`
}
