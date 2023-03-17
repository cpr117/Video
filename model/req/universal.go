// @User CPR
package req

// 关键字查询
type KeywordQuery struct {
	Keyword string `form:"keyword"`
}

// 获取数据(需要分页)
type PageQuery struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	KeyWord  string `form:"keyword"`
}
