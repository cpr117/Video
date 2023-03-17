// @User CPR
package service

import (
	"VideoWeb/dao"
	"VideoWeb/model"
	"VideoWeb/model/req"
	"VideoWeb/model/resp"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"strconv"
	"time"
)

type Comment struct{}

//func (*Comment) GetFrontList(req req.GetFrontComments) (ret resp.PageResult[[]resp.FrontCommentVo]) {
//	// 数据库查询评论列表
//	rootCommentList, total := commentDao.GetFrontList(req)
//	commentIds := make([]int, 0)
//	// [文章ID : 对应点赞数]
//	likeCountMap := utils.Redis.HGetAll(KEY_COMMENT_LIKE_COUNT)
//	dislikeCountMap := utils.Redis.HGetAll(KEY_COMMENT_DISLIKE_COUNT)
//
//	for _, comment := range rootCommentList {
//		//rootCommentList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(comment.ID)])
//		//rootCommentList[i].DislikeCount, _ = strconv.Atoi(dislikeCountMap[strconv.Itoa(comment.ID)])
//		if comment.ReplyCommentId == 0 {
//			commentIds = append(commentIds, comment.ID)
//		}
//	}
//
//	// 获取回复列表 包含根评论
//	commentList := commentDao.GetReplyList(commentIds)
//
//	commentMap := make(map[int]*resp.FrontCommentVo)
//	commentCountMap := make(map[int]int)
//
//	// 首先处理 点赞 和 踩
//	// 其次将frontCommentVo以地址的形式丢到commentMap里面方便后续操作
//	for i, reply := range commentList {
//		commentList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(reply.ID)])
//		commentList[i].DislikeCount, _ = strconv.Atoi(dislikeCountMap[strconv.Itoa(reply.ID)])
//		commentMap[reply.ID] = &reply
//		_, ok := commentCountMap[reply.ReplyCommentId]
//		if ok {
//			commentCountMap[reply.ReplyCommentId]++
//		} else {
//			commentCountMap[reply.ReplyCommentId] = 1
//		}
//	}
//
//	// 理论上是判断如果没有后继节点即ReplyCount=0则将整个对象添加到前驱节点的ReplyVoList中 同时将后继节点的ReplyCount--
//	for i := 0; i < len(commentList); i++ {
//		j := i
//		if _, ok := commentCountMap[commentList[j].ID]; !ok { // 如果没有回复，则将其添加到父评论的回复列表中
//			commentMap[commentList[j].ReplyCommentId].ReplyVoList =
//				append(commentMap[commentList[j].ReplyCommentId].ReplyVoList, *commentMap[j])
//
//			commentCountMap[commentList[j].ReplyCommentId]--
//			for commentCountMap[commentList[j].ReplyCommentId] == 0 {
//				// 如果父评论只有一个回复，则将其从commentCountMap中删除 因为唯一一个回复已经添加到父评论的回复列表中
//				delete(commentCountMap, commentList[j].ID)
//				j = commentList[j].ReplyCommentId
//				commentMap[commentList[j].ReplyCommentId].ReplyVoList =
//					append(commentMap[commentList[j].ReplyCommentId].ReplyVoList, commentList[j])
//				commentCountMap[commentList[j].ReplyCommentId]--
//			}
//		} else { // 如果有回复，则将其添加到父评论的回复列表中
//			commentCountMap[commentList[j].ReplyCommentId]--
//		}
//	}
//
//	//for i, comment := range commentList {
//	//	commentList[i].ReplyVoList = commentMap[comment.ID].ReplyVoList
//	//}
//	ret.Total = total
//	ret.List =
//
//	return
//}

func (*Comment) GetFrontList(req req.GetFrontComments) (ret resp.PageResult[[]resp.FrontCommentVo]) {
	// 数据库查询评论列表
	rootCommentList, total := commentDao.GetFrontList(req)
	commentIds := make([]int, 0)
	// [文章ID : 对应点赞数]
	likeCountMap := utils.Redis.HGetAll(KEY_COMMENT_LIKE_COUNT)
	dislikeCountMap := utils.Redis.HGetAll(KEY_COMMENT_DISLIKE_COUNT)

	for _, comment := range rootCommentList {
		//rootCommentList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(comment.ID)])
		//rootCommentList[i].DislikeCount, _ = strconv.Atoi(dislikeCountMap[strconv.Itoa(comment.ID)])
		if comment.ReplyCommentId == 0 {
			commentIds = append(commentIds, comment.ID)
		}
	}

	// 获取回复列表 包含根评论
	cList := commentDao.GetReplyList(commentIds)

	//cMap := make(map[int]*resp.FrontCommentVo)
	cCountMap := make(map[int]int)
	CId := make(map[int]int) // KEY:id VALUE:index

	// 首先处理 点赞 和 踩
	// 其次将frontCommentVo以地址的形式丢到commentMap里面方便后续操作
	for i, reply := range cList {
		cList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(reply.ID)])
		cList[i].DislikeCount, _ = strconv.Atoi(dislikeCountMap[strconv.Itoa(reply.ID)])
		CId[reply.ID] = i
		//cMap[reply.ID] = &reply
		_, ok := cCountMap[reply.ReplyCommentId]
		if ok {
			cCountMap[reply.ReplyCommentId]++
		} else {
			cCountMap[reply.ReplyCommentId] = 1
		}
	}
	list := make([]resp.FrontCommentVo, 0)
	//for range循环遍历slice/map，值是复制的，且每次循环都是用同一个值保存复制后的值
	for i, comment := range cList {
		if _, ok := cCountMap[comment.ID]; !ok && comment.ReplyCommentId == 0 { // 说明即是根节点也是子节点
			list = append(list, cList[i])
			//list = append(list, comment) 不知道有没有区别
			continue
		}
		if _, ok := cCountMap[comment.ID]; !ok { // 表示该评论为子评论 只处理子评论 因为子评论依据拓扑结构向上进行处理
			cList[CId[comment.ReplyCommentId]].ReplyVoList = append(cList[CId[comment.ReplyCommentId]].ReplyVoList, cList[i])
			cCountMap[comment.ReplyCommentId]--

			cId := comment.ReplyCommentId
			cIndex := CId[cId]
			for cCountMap[cId] == 0 { // 父节点的子节点全部处理完毕 进行下一步
				if cList[cIndex].ReplyCommentId == 0 { // 说明已经是根节点
					list = append(list, cList[cIndex])
					break
				}
				delete(cCountMap, cId) // j是key 不用担心乱序的问题 因为是set不是slice

				cList[CId[cList[cIndex].ReplyCommentId]].ReplyVoList =
					append(cList[CId[cList[cIndex].ReplyCommentId]].ReplyVoList, cList[cIndex])
				cId = cList[cIndex].ReplyCommentId
				cCountMap[cId]--
			}

			//j := comment.ReplyCommentId
			//for cCountMap[j] == 0 { // 父节点的子节点全部处理完毕 进行下一步
			//	if cList[CId[j]].ReplyCommentId == 0 { // 说明已经是根节点
			//		list = append(list, cList[CId[j]])
			//		break
			//	}
			//	delete(cCountMap, j) // j是key 不用担心乱序的问题 因为是set不是slice
			//
			//	cList[CId[cList[CId[j]].ReplyCommentId]].ReplyVoList =
			//		append(cList[CId[cList[CId[j]].ReplyCommentId]].ReplyVoList, cList[CId[j]])
			//	j = cList[CId[j]].ReplyCommentId
			//	cCountMap[j]--
			//}
		}
	}
	ret.List = list
	ret.Total = total
	ret.PageSize = req.PageSize
	ret.PageNum = req.PageNum
	return
}

func (*Comment) SaveComment(req req.SaveComment) int {
	comment := utils.CopyProperties[model.Comment](req) // vo -> po
	comment.CreatedAt = time.Now()
	comment.IsReview = BiliInfoService.GetBiliConfig().IsCommentReview
	dao.Create(&comment)
	return r.OK
}

func (*Comment) SoftDeleteComment(req req.DeleteComment) int {
	if checkUserDeleteAuth(req) {
		data := dao.GetOne[model.Comment]("id = ?", req.CommentId)
		dao.Update[model.Comment](&data, "is_delete = 1")
		return r.OK
	} else {
		return r.ERROR_COMMENT_DELETE_AUTH
	}
}

func (*Comment) HardDeleteComment(req req.DeleteComment) int {
	if checkUserDeleteAuth(req) {
		dao.Delete[model.Comment]("comment_id = ?", req.CommentId)
		return r.OK
	} else {
		return r.ERROR_COMMENT_DELETE_AUTH
	}
}

func checkUserDeleteAuth(req req.DeleteComment) bool {
	return dao.Count[model.Comment]("comment_id = ? and user_id = ?", req.CommentId, req.UserId) == 1
}
