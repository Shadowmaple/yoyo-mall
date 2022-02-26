package collection

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/collection"

	"github.com/gin-gonic/gin"
)

type DeleteReq struct {
	List        []uint32 `json:"list"`       // 记录id列表，批量删除，优先
	ProductList uint32   `json:"product_id"` // 商品id，根据用户id和商品id删除
}

func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover error:", err)
		}
	}()

	req := &DeleteReq{}
	if err := c.BindJSON(req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("id").(uint32)

	var err error
	if len(req.List) != 0 {
		err = collection.BatchDelete(userID, req.List)
	} else {
		err = collection.DelByProductID(userID, req.ProductList)
	}
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
