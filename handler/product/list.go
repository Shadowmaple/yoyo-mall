package product

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

type ListReq struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Cid   uint32 `form:"cid"`
	Cid2  uint32 `form:"cid2"`
	Sort  int    `form:"sort"`
}

type ListResp struct {
	Total int                    `json:"total"`
	List  []*product.ProductItem `json:"list"`
}

func List(c *gin.Context) {
	req := &ListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	if req.Limit == 0 {
		req.Limit = 20
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
