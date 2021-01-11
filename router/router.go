package router

import (
	"pku-class/market/handler/manage"
	"pku-class/market/handler/menu"
	"pku-class/market/handler/shop"
	"pku-class/market/handler/user"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	router.Use(JWY())

	manageGroup := router.Group("/manage")
	shopGroup := router.Group("/shop")
	userGroup := router.Group("/user")
	menuGroup := router.Group("/menu")
	{
		manageGroup.POST("/add/", manage.AddPostHandler)
		manageGroup.POST("/delete/", manage.DeletePostHandler)
		manageGroup.POST("/update/", manage.UpdatePostHandler)
	}
	{
		shopGroup.POST("/flash/", shop.FlashPostHandler)
		shopGroup.POST("/consume/", shop.ConsumePostHandler)
		shopGroup.POST("/", shop.PostHandler)
	}
	{
		menuGroup.POST("/", menu.PostHandler)
	}
	{
		userGroup.POST("/charge/", user.ChargePostHandler)
		userGroup.POST("/order/delete/", user.OrderDeletePostHandler)
		userGroup.GET("/order/", user.OrderGetHandler)
		userGroup.POST("/login/", user.LoginPostHandler)
	}
	return router
}
