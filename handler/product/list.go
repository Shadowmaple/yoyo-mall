package product

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

type ListReq struct {
	Limit int
	Page  int
	Cid   uint32
	Cid2  uint32
	Sort  int
}

type ListResp struct {
	Total int
	List  []*product.ProductItem
}

func List(c *gin.Context) {
	req := &ListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	var userID uint32
	idTemp, ok := c.Get("id")
	if ok {
		userID = idTemp.(uint32)
	}

	filter := product.FilterItem{
		Cid:  req.Cid,
		Cid2: req.Cid2,
		Sort: req.Sort,
	}
	list, err := product.List(userID, req.Limit, req.Page, filter)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, ListResp{
		Total: len(list),
		List:  list,
	})
}
