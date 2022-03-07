package router

import (
	"net/http"
	"yoyo-mall/handler/address"
	"yoyo-mall/handler/cart"
	"yoyo-mall/handler/category"
	"yoyo-mall/handler/collection"
	"yoyo-mall/handler/comment"
	"yoyo-mall/handler/coupon"
	"yoyo-mall/handler/feedback"
	"yoyo-mall/handler/logistics"
	"yoyo-mall/handler/order"
	"yoyo-mall/handler/product"
	"yoyo-mall/handler/search"
	"yoyo-mall/handler/upload"
	"yoyo-mall/handler/user"
	"yoyo-mall/router/middleware"

	"github.com/gin-gonic/gin"
)

const BasePath = "/api/v1"

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(gin.Logger())
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	authFunc := middleware.AuthMiddleware
	adminAuthFunc := middleware.AuthMiddleware
	visitorAuthFunc := middleware.VisitorAuthMiddleware

	// auth
	g.POST(BasePath+"/login", user.Login)
	g.POST(BasePath+"/login/admin", user.AdminLogin)

	// upload
	g.POST(BasePath+"/upload/image", authFunc(), upload.Image)

	// user
	u := g.Group(BasePath + "/user")
	u.Use(authFunc())
	{
		u.GET("/info", user.GetInfo)
		u.POST("/info", user.UpdateInfo)
	}

	// product
	p := g.Group(BasePath + "/product")
	{
		p.GET("/list", visitorAuthFunc(), product.List)
		p.GET("/info/:id", visitorAuthFunc(), product.GetInfo)
		p.GET("/rank", visitorAuthFunc(), product.GetRank)
		p.POST("", adminAuthFunc(), product.CreateOrUpdate)
		p.DELETE("/:id", adminAuthFunc(), product.Delete)
	}

	// category
	cate := g.Group(BasePath + "/category")
	cate.GET("", category.Get)

	// cart 购物车
	cartGp := g.Group(BasePath + "/cart")
	cartGp.Use(authFunc())
	{
		cartGp.GET("", cart.Get)
		cartGp.POST("", cart.Add)
		cartGp.DELETE("", cart.Delete)
		cartGp.PUT("", cart.Update)
	}

	// order 订单
	orderGp := g.Group(BasePath + "/order")
	orderGp.Use(authFunc())
	{
		orderGp.GET("/list", order.List)
		orderGp.GET("/info/:id", order.Info)
		orderGp.POST("", order.Create)
		orderGp.PUT("/:id", order.Update)
		orderGp.GET("/admin/list", order.AdminList)
	}

	// search
	searchGp := g.Group(BasePath + "/search")
	{
		searchGp.GET("/product", visitorAuthFunc(), search.ProductSearch)
		searchGp.GET("/order", authFunc(), search.OrderSearch)
	}

	// collection 收藏
	collectionGroup := g.Group(BasePath + "/collection")
	collectionGroup.Use(authFunc())
	{
		collectionGroup.GET("", collection.List)
		collectionGroup.POST("", collection.Add)
		collectionGroup.DELETE("", collection.Delete)
	}

	// logistic 物流
	logisticsGp := g.Group(BasePath + "/logistics")
	{
		logisticsGp.GET("", authFunc(), logistics.Get)
		logisticsGp.GET("/list", adminAuthFunc(), logistics.List)
		logisticsGp.PUT("", adminAuthFunc(), logistics.Update)
	}

	// address
	addrGp := g.Group(BasePath + "/address")
	addrGp.Use(authFunc())
	{
		addrGp.GET("", address.Get)
		addrGp.POST("", address.AddOrUpdate)
		addrGp.DELETE("", address.Delete)
	}

	// evaluation, comment
	commentGp := g.Group(BasePath + "/evaluation")
	{
		commentGp.GET("/list", visitorAuthFunc(), comment.EvaluationList)
		commentGp.POST("", authFunc(), comment.EvaluationCreateOrUpdate)
		commentGp.GET("/info/:id", visitorAuthFunc(), comment.EvaluationInfo)
		commentGp.GET("/info/:id/comment", visitorAuthFunc(), comment.CommentList)
		commentGp.POST("/info/:id/comment", authFunc(), comment.CommentCreateOrUpdate)
	}

	// like
	g.POST(BasePath+"/like", authFunc(), comment.Like)

	// coupon 优惠券
	couponGp := g.Group(BasePath + "/coupon")
	{
		couponGp.GET("/private", authFunc(), coupon.PrivateList)
		couponGp.GET("/public", authFunc(), coupon.PublicList)
		couponGp.GET("", authFunc(), coupon.Grab)
		couponGp.PUT("/private", authFunc(), coupon.Use)
		couponGp.POST("", adminAuthFunc(), coupon.AddOrUpdate)
		couponGp.DELETE("", adminAuthFunc(), coupon.Delete)
		couponGp.GET("/admin", adminAuthFunc(), coupon.AdminList)
	}

	// feedback
	feedbackGp := g.Group(BasePath + "/feedback")
	{
		feedbackGp.POST("", authFunc(), feedback.Add)
		feedbackGp.GET("", adminAuthFunc(), feedback.List)
		feedbackGp.POST("/read", adminAuthFunc(), feedback.Read)
	}

	// message
	// ...

	return g
}
