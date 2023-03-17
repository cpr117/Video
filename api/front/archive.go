// @User CPR
package front

import (
	"VideoWeb/model/req"
	"VideoWeb/utils"
	"VideoWeb/utils/r"
	"github.com/gin-gonic/gin"
)

type Collection struct{}

func (*Collection) GetCollectionList(c *gin.Context) {
	r.SuccessData(c, collectionService.GetCollection(utils.BindQuery[req.GetCollection](c)))
}

func (*Collection) AddToCollection(c *gin.Context) {
	r.SendCode(c, collectionService.AddCollection(utils.BindJson[req.AddToCollection](c)))
}

func (*Collection) CreateCollection(c *gin.Context) {
	r.SendCode(c, collectionService.CreateCollection(utils.BindValidJson[req.CreateCollection](c)))
}

func (*Collection) UpdateCollection(c *gin.Context) {
	r.SendCode(c, collectionService.UpdateCollection(utils.BindValidJson[req.CreateCollection](c)))
}

func (*Collection) DeleteCollection(c *gin.Context) {
	r.SendCode(c, collectionService.DeleteCollection(utils.GetIntParam(c, "id")))
}

func (*Collection) DelFromCollection(c *gin.Context) {
	r.SendCode(c, collectionService.DelFromCollection(utils.BindJson[req.DelFromCollection](c)))
}
