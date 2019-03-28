package router

import (
	"cladmin/handler/menu"
	"cladmin/handler/role"
	"cladmin/handler/sd"
	"cladmin/handler/upload"
	"cladmin/handler/user"
	"cladmin/router/middleware"
	"cladmin/router/middleware/permission"
	"github.com/gin-gonic/gin"
	"net/http"
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
	//api get AliyunOss signature
	g.GET("/oss/generatesignature", sd.GenerateSignature)

	apiV1 := g.Group("/v1")
	apiV1.Use(middleware.AuthMiddleware())
	apiV1.Use(permission.CasbinMiddleware())
	{
		apiV1.POST("/users/create", user.Create)
		apiV1.PUT("/users/update", user.Update)
		apiV1.GET("/users/get", user.Get)
		apiV1.GET("/users/list", user.List)
		apiV1.DELETE("/users/del", user.Delete)

		apiV1.POST("/roles/create", role.Create)
		apiV1.GET("/roles/get", role.Get)
		apiV1.GET("/roles/list", role.List)
		apiV1.PUT("/roles/update", role.Update)
		apiV1.DELETE("/roles/del", role.Delete)

		apiV1.POST("/menus/create", menu.Create)
		apiV1.PUT("/menus/update", menu.Update)
	}
	//user
	/*userRouter := g.Group("/v1/user")
	{
		// Need auth
		userRouter.GET("", middleware.AuthMiddleware(), user.Get)
		userRouter.PUT("", middleware.AuthMiddleware(), user.Update)
		//No Need auth
		userRouter.POST("", user.Create)
		userRouter.DELETE("/:id", user.Delete)
		userRouter.GET("/index", user.List)
	}*/
	//article
	/*articleRouter := g.Group("/v1/article")
	{
		articleRouter.POST("", article.Create)
		articleRouter.GET("/:id", article.Get)
		articleRouter.GET("", article.List)
	}*/
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
