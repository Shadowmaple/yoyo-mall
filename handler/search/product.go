package search

import (
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/product"

	"github.com/gin-gonic/gin"
)

type ProductSearchReq struct {
	Limit     int    `form:"limit"`
	Page      int    `form:"page"`
	Title     string `form:"title"`
	Book      string `form:"book"`
	Author    string `form:"author"`
	Publisher string `form:"publisher"`
}

func ProductSearch(c *gin.Context) {
	req := &ProductSearchReq{}
	if err := c.BindQuery(req); err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}

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
