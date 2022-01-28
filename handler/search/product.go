package search

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

type ProductSearchReq struct {
	Limit     int
	Page      int
	Title     string
	Book      string
	Author    string
	Publisher string
}

func ProductSearch(c *gin.Context) {
	req := &ProductSearchReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	log.Printf("%+v", req)

	userID := c.GetUint("id")

	searchItem := product.SearchItem{
		Title:     req.Title,
		Book:      req.Book,
		Author:    req.Author,
		Publisher: req.Publisher,
	}

	list, err := product.Search(uint32(userID), req.Limit, req.Page, searchItem)
	if err != nil {
		handler.SendError(c, errno.InternalError, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, handler.ListResp{
		Total: len(list),
		List:  list,
	})
}
