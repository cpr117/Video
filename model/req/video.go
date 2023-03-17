// @User CPR
package req

type SaveVideo struct {
	Title       string `json:"title"`
	Desc        string `gorm:"not null;comment:视频描述" json:"desc"`
	PartitionId int    `gorm:"int(11);not null;comment:视频分类id" json:"partition_id"`
	AuthorId    int    `gorm:"int(11);comment:视频作者id" json:"author_id"`
}

//{"video_name":"cpr","video_desc":"geebar","video_type":"日常视频","video_author_id":1}

type GetVideos struct {
	PageQuery
	VideoName  string `json:"video_name"`
	VideoParId int    `json:"video_par_id"`
}

type UpdateVideo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Cover       string `json:"cover"`
	Desc        string `json:"desc"`
	PartitionId int    `json:"partition_id"`
	Status      int8   `json:"status"`
}
