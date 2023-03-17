// @User CPR
package service

import (
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
)

type Category struct{}

func (*Category) GetFrontList() (ret resp.ListResult[[]resp.CategoryVo]) {
	data, total := categoryDao.GetCategoryList(req.PageQuery{})
	ret.List = data
	ret.Total = total
	return ret
}
