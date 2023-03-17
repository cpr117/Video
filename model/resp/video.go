// @User CPR
package resp

// 视频列表
type VideoVo struct {
	Id            int    `json:"id"`             // 视频id
	Title         string `json:"title"`          // 视频标题
	PartitionId   int    `json:"partition_id"`   // 分区id
	PartitionName string `json:"partition_name"` // 分区名称
	Desc          string `json:"video_desc"`     // 视频描述
	Cover         string `json:"cover"`          // 视频封面

	AuthorId int `json:"author_id"` // 视频作者id
	//Articulation string `json:"articulation"` // 视频清晰度
	Url string `json:"video_url"` // 视频地址
	//Avatar string `json:"avatar"`		// 暂时还没有保存图片的部分
}
