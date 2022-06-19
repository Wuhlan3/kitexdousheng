package main

import (
	// "dousheng/controller"
	// "dousheng/middleware"
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/handlers"
	"kitexdousheng/pkg/middleware"
	//"path"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", middleware.AuthMiddleware(), handlers.Feed)
	apiRouter.GET("/user/", middleware.AuthMiddleware(), handlers.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	apiRouter.POST("/publish/action/", middleware.AuthMiddleware(), handlers.PublishAction)
	apiRouter.GET("/publish/list/", middleware.AuthMiddleware(), handlers.PublishList)

	// // extra apis - I
	// apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	// apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.AuthMiddleware(), handlers.CommentAction)
	apiRouter.GET("/comment/list/", middleware.AuthMiddleware(), handlers.CommentList)

	// // extra apis - II
	apiRouter.POST("/relation/action/", handlers.RelationAction)
	apiRouter.GET("/relation/follow/list/", handlers.FollowList)
	apiRouter.GET("/relation/follower/list/", handlers.FollowerList)
}
