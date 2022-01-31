package category

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/category"

	"github.com/gin-gonic/gin"
)

type GetResp struct {
	Total int                 `json:"total"`
	List  []*category.CidItem `json:"list"`
}

func Get(c *gin.Context) {
	list, err := category.GetList()
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, GetResp{
		Total: len(list),
		List:  list,
	})
}
