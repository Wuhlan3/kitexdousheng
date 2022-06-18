package main

import (
	// "dousheng/controller"
	// "dousheng/middleware"
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/handlers"
	//"path"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	//r.Static("/static", "./public")
	// r.GET("/home/go/src/dousheng/public/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	filename := path.Join("./public/", name)
	// 	c.File(filename)
	// 	return
	// })

	apiRouter := r.Group("/douyin")

	// basic apis
	// apiRouter.GET("/feed/", controller.Feed)
	// apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	// apiRouter.POST("/publish/action/", controller.Publish)
	// apiRouter.GET("/publish/list/", controller.PublishList)

	// // extra apis - I
	// apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	// apiRouter.GET("/favorite/list/", controller.FavoriteList)
	// apiRouter.POST("/comment/action/", controller.CommentAction)
	// apiRouter.GET("/comment/list/", controller.CommentList)

	// // extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction)
	// apiRouter.GET("/relation/follow/list/", controller.FollowList)
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
