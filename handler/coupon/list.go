package coupon

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/coupon"

	"github.com/gin-gonic/gin"
)

type PrivateListReq struct {
	Status int8 `form:"status"`
}

type PublicListReq struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Cid   uint32 `form:"cid"`
	Cid2  uint32 `form:"cid2"`
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

	if req.Limit <= 0 {
		req.Limit = 20
	}

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

type AdminListReq struct {
	Kind  int8   `form:"kind"` // 0全部……
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Cid   uint32 `form:"cid"`
	Cid2  uint32 `form:"cid2"`
}

func AdminList(c *gin.Context) {
	req := &AdminListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}

	list, err := coupon.AdminList(req.Page, req.Limit, req.Cid, req.Cid2, req.Kind)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
