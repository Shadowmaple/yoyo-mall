package comment

import (
	"strconv"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/comment"

	"github.com/gin-gonic/gin"
)

type CommentListReq struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

func CommentList(c *gin.Context) {
	evaluationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	req := &CommentListReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userID := c.GetUint("id")

	list, err := comment.List(uint32(userID), uint32(evaluationID), req.Limit, req.Page)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}

func CommentCreateOrUpdate(c *gin.Context) {
	evaluationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	req := &comment.BasicItem{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	if req.ID == 0 {
		err = comment.Publish(userID, uint32(evaluationID), req)
	} else {
		err = comment.Update(userID, uint32(evaluationID), req)
	}
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
