package comment

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"

	"github.com/gin-gonic/gin"
)

type LikeReq struct {
	ID             uint32 `json:"id"`
	Kind           int8   `json:"kind"`
	ExpectedStatus bool   `json:"expected_status"`
}

type LikeResp struct {
	HasLiked bool `json:"has_liked"`
}

func Like(c *gin.Context) {
	req := &LikeReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	var err error
	if req.ExpectedStatus {
		err = model.Like(userID, req.ID, req.Kind)
	} else {
		err = model.Unlike(userID, req.ID, req.Kind)
	}

	if err != nil {
		log.Printf("like api: req: %+v\n", req)
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, LikeResp{
		HasLiked: req.ExpectedStatus,
	})
}
