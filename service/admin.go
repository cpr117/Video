// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils/r"
)

type Admin struct{}

func (*Admin) GetList(req req.GetAdminList) (ret resp.PageResult[[]resp.AdminList]) {
	list, total := adminDao.GetList(req)
	ret.List = list
	ret.Total = total
	ret.PageSize = req.PageSize
	ret.PageNum = req.PageNum
	return
}

func (*Admin) Create(req req.CreateAdmin) int {
	newAdmin := &model.Admin{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     2, // 默认是2
		Level:    req.Level,
	}
	if req.Role != 0 {
		newAdmin.Role = req.Role
	}
	dao.Create(&newAdmin)
	return r.OK
}

func (Admin) Update(req req.UpdateAdmin) int {
	dao.Create(&model.Admin{
		Universal: model.Universal{Id: req.Id},
		Password:  req.Password,
		Email:     req.Email,
		Phone:     req.Phone,
	})
	return r.OK
}
