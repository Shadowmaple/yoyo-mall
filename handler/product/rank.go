package product

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

type RankReq struct {
	Kind int
	Cid  uint32
	Cid2 uint32
}

// 获取榜单
func GetRank(c *gin.Context) {
	req := &RankReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	list, err := product.GetRank(req.Kind, 20, req.Cid, req.Cid2)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, map[string]interface{}{
		"total": len(list),
		"list":  list,
	})
}
