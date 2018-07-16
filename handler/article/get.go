package article

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	"strconv"
	"apiserver/pkg/errno"
	. "apiserver/handler"
	"github.com/json-iterator/go"
)

func Get(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article, err := model.GetArticle(uint64(articleId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// unmarshal json fields
	images := jsonParse(article)
	rep := &CreateResponse{
		Id:        article.Id,
		CateId:    article.CateId,
		Title:     article.Title,
		Images:    images,
		Author:    Author{Id: article.Author.Id, Username: article.Author.Username},
		Content:   article.Content,
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	SendResponse(c, nil, rep)
}

func jsonParse(a *model.ArticleModel) (images []string) {
	jsoniter.Unmarshal([]byte(a.Images), &images)
	return
}
