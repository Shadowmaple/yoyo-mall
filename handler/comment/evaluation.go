package comment

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/evaluation"

	"github.com/gin-gonic/gin"
)

type EvaluationListReq struct {
	Limit     int    `form:"limit"`
	Page      int    `form:"page"`
	ProductID uint32 `form:"product_id"`
}

func EvaluationList(c *gin.Context) {
	req := &EvaluationListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userID := c.GetUint("id")

	list, err := evaluation.List(uint32(userID), req.ProductID, req.Limit, req.Page)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}

func EvaluationCreateOrUpdate(c *gin.Context) {
	req := &evaluation.BasicItem{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	var err error
	if req.ID == 0 {
		err = evaluation.New(userID, req)
	} else {
		err = evaluation.Update(userID, req)
	}

	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
