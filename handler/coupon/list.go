package coupon

import (
	"fmt"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/coupon"

	"github.com/gin-gonic/gin"
)

type PrivateListReq struct {
	Status int8
}

type PublicListReq struct {
	Page  int
	Limit int
	Cid   uint32
	Cid2  uint32
}

func PrivateList(c *gin.Context) {
	req := &PrivateListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	list, err := coupon.PrivateList(userID, req.Status)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}

func PublicList(c *gin.Context) {
	req := &PublicListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	fmt.Printf("%+v", req)

	userID := c.MustGet("id").(uint32)

	list, err := coupon.PublicList(userID, req.Page, req.Limit, req.Cid, req.Cid2)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
