package router

import (
	"apiserver/handler/sd"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"apiserver/handler/article"
	"apiserver/handler/user"
	"apiserver/handler/upload"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// api for authentication functionalities
	g.POST("/login", user.Login)
	g.GET("/refresh", user.Refresh)
	//user
	userRouter := g.Group("/v1/user")
	{
		// Need auth
		userRouter.GET("", middleware.AuthMiddleware(), user.Get)
		userRouter.PUT("", middleware.AuthMiddleware(), user.Update)
		//No Need auth
		userRouter.POST("", user.Create)
		userRouter.DELETE("/:id", user.Delete)
		userRouter.GET("/index", user.List)
	}
	//article
	articleRouter := g.Group("/v1/article")
	{
		articleRouter.POST("", article.Create)
		articleRouter.GET("/:id", article.Get)
		articleRouter.GET("", article.List)
	}
	//upload
	uploadRouter := g.Group("/v1/upload")
	{
		uploadRouter.POST("/image", upload.Img)
	}
	//The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		//svcd.GET("/demo1", sd.DemoOne)
	}
	//public static
	publicRouter := g.Group("/public")
	{
		publicRouter.Static("", "public")
	}
	return g
}
