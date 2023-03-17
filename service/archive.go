// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
)

type Collection struct{}

func (*Collection) GetCollection(req req.GetCollection) (ret resp.PageResult[[]resp.VideoVo]) {
	if req.UserId == 0 || req.CollectionId == 0 {
		utils.Logger.Error("用户ID或收藏夹ID不能为空")
		return
	}
	videoList, total := archiveDao.GetCollectionListByUserId(req)
	ret.List = videoList
	ret.Total = total
	ret.PageNum = req.PageNum
	ret.PageSize = req.PageSize
	return
}

func (*Collection) AddCollection(req req.AddToCollection) int {
	if req.UserId == 0 || req.CollectionId == 0 || req.VideoId == 0 {
		utils.Logger.Error("用户ID或收藏夹ID或视频ID不能为空")
		return r.ERROR_UPDATE_COLLECTION
	}
	if dao.GetOne[model.UserCollection]("user_collection", "user_id = ? and collection_id = ?", req.UserId, req.CollectionId).UserId != 0 {
		data := model.VideoCollection{
			VideoId:      req.VideoId,
			CollectionId: req.CollectionId,
		}
		dao.Create(&data)
		return r.OK
	} else {
		utils.Logger.Info("用户未创建该收藏夹")
		return r.ERROR_COLLECTION_NOT_EXIST
	}

}

func (*Collection) CreateCollection(req req.CreateCollection) int {
	if req.UserId == 0 || req.CollectionId == 0 {
		utils.Logger.Error("用户ID或收藏夹名不能为空")
		return r.ERROR_UPDATE_COLLECTION
	}
	data := model.Collection{
		UserId:    req.UserId,
		Name:      req.CollectionName,
		Count:     0,
		IsPrivacy: req.IsPrivacy,
	}
	dao.Create(&data)

	uc := model.UserCollection{
		UserId:       req.UserId,
		CollectionId: data.Id,
	}
	dao.Create(&uc)
	return r.OK
}

func (*Collection) UpdateCollection(req req.CreateCollection) int {
	// 查询是否存在同名收藏夹
	existByName := dao.GetOne[model.Collection]("collection", "name = ?", req.CollectionName)
	// (同名存在) && (存在的id不等于当前要更新的id) -> 重复
	if !existByName.IsEmpty() && existByName.Id != req.CollectionId {
		utils.Logger.Info("收藏夹名重复")
		return r.ERROR_COLLECTION_NAME_REPEAT
	}
	collection := utils.CopyProperties[model.Collection](req) // vo->po

	if req.CollectionId != 0 {
		dao.Update(&collection, "name")
	} else {
		dao.Create(&collection)
	}
	return r.OK
}

func (*Collection) DeleteCollection(id int) (code int) {
	count := dao.Count[model.VideoCollection]("collection_id = ?", id)
	if count > 0 {
		dao.Delete[model.VideoCollection]("collection_id = ?", id)
	}
	dao.Delete[model.Collection]("id = ?", id)
	return r.OK
}

func (*Collection) DelFromCollection(req req.DelFromCollection) (code int) {
	count := dao.Count[model.VideoCollection]("collection_id = ? and video_id = ?", req.CollectionId, req.VideoId)
	if count == 0 {
		return r.ERROR_COLLECTION_VIDEO_MISMATCH
	}
	dao.Delete[model.VideoCollection]("collection_id = ? and video_id = ?", req.CollectionId, req.VideoId)
	return r.OK
}
