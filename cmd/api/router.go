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
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)

	authRouter := apiRouter.Group("")
	authRouter.Use(middleware.AuthMiddleware())
	{
		authRouter.GET("/feed/", handlers.Feed)
		authRouter.POST("/relation/action/", handlers.RelationAction)
		authRouter.GET("/user/", handlers.UserInfo)
		authRouter.POST("/publish/action/", handlers.PublishAction)
		authRouter.GET("/publish/list/", handlers.PublishList)
		authRouter.GET("/relation/follow/list/", handlers.FollowList)
		authRouter.GET("/relation/follower/list/", handlers.FollowerList)
		authRouter.POST("/comment/action/", handlers.CommentAction)
		authRouter.GET("/comment/list/", handlers.CommentList)
		authRouter.POST("/favorite/action/", handlers.FavoriteAction)
		authRouter.GET("/favorite/list/", handlers.FavoriteList)
	}

	// // extra apis - I
	// apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	// apiRouter.GET("/favorite/list/", controller.FavoriteList)

	// // extra apis - II

}
