package main

import (
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "../../public")

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", handlers.DouyinFeed)
		// apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", handlers.UserRegister)
		// apiRouter.POST("/user/login/", controller.Login)
		// apiRouter.POST("/publish/action/", controller.Publish)
		// apiRouter.GET("/publish/list/", controller.PublishList)
	}

	auth := r.Group("/douyin")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/user", handlers.GetUserInfoById)
	}

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
