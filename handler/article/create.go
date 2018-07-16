package article

import (
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/model"
	"github.com/json-iterator/go"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//json fields handler and bind articleModel
	a := bindArticle(r)

	//validate
	if err := a.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//create
	if err := a.CreateArticle(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	rep := CreateResponse{
		Id:    a.Id,
		Title: a.Title,
	}
	SendResponse(c, nil, rep)
}

func bindArticle(r CreateRequest) (a model.ArticleModel) {
	imgJson, _ := jsoniter.Marshal(r.Images)
	a.Uid = r.Uid
	a.CateId = r.CateId
	a.Title = r.Title
	a.Images = string(imgJson)
	return
}
