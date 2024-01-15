package router

import (
	"cladmin/handler/aliyun/oss"
	"cladmin/handler/aliyun/sts"
	"cladmin/handler/article"
	"cladmin/handler/bulletin"
	"cladmin/handler/category"
	"cladmin/handler/config"
	"cladmin/handler/dept"
	"cladmin/handler/dict"
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
	g.GET("/logout", middleware.AuthMiddleware(), user.Logout)
	g.GET("/refresh", user.Refresh)
	//api get AliyunOss signature
	g.GET("/oss/generatesignature", oss.WebUploadSign)
	g.GET("/sts/assumeRole", sts.GetAssumeRole)

	apiV1 := g.Group("/v1")
	apiV1.Use(middleware.AuthMiddleware())
	apiV1.Use(permission.CasbinMiddleware())
	{
		apiV1.POST("/users/create", user.Create)
		apiV1.PUT("/users/update", user.Update)
		apiV1.GET("/users/get", user.Get)
		apiV1.GET("/users/list", user.List)
		apiV1.DELETE("/users/delete", user.Delete)
		apiV1.GET("/users/personal", user.GetPersonalInfo)
		apiV1.PUT("/users/updatePersonal", user.UpdatePersonal)
		apiV1.POST("/users/logout", user.LogoutUser)

		apiV1.POST("/roles/create", role.Create)
		apiV1.GET("/roles/get", role.Get)
		apiV1.GET("/roles/list", role.List)
		apiV1.PUT("/roles/update", role.Update)
		apiV1.DELETE("/roles/delete", role.Delete)
		apiV1.GET("/roles/select", role.Select)

		apiV1.POST("/menus/create", menu.Create)
		apiV1.PUT("/menus/update", menu.Update)
		apiV1.GET("/menus/get", menu.Get)
		apiV1.GET("/menus/list", menu.List)
		apiV1.DELETE("/menus/delete", menu.Delete)
		apiV1.GET("/menus/nav", menu.GetMenuNav)
		apiV1.GET("/menus/permissions", menu.GetPermissions)

		apiV1.GET("/config/get", config.Get)
		apiV1.POST("/config/create", config.Create)
		apiV1.PUT("/config/update", config.Update)
		apiV1.GET("/config/list", config.List)
		apiV1.DELETE("/config/delete", config.Delete)

		apiV1.POST("/categories/create", category.Create)
		apiV1.PUT("/categories/update", category.Update)
		apiV1.GET("/categories/get", category.Get)
		apiV1.GET("/categories/list", category.List)
		apiV1.DELETE("/categories/delete", category.Delete)

		apiV1.POST("/articles/create", article.Create)
		apiV1.PUT("/articles/update", article.Update)
		apiV1.GET("/articles/get", article.Get)
		apiV1.GET("/articles/list", article.List)
		apiV1.DELETE("/articles/delete", article.Delete)

		apiV1.POST("/bulletin/create", bulletin.Create)
		apiV1.PUT("/bulletin/update", bulletin.Update)
		apiV1.GET("/bulletin/get", bulletin.Get)
		apiV1.GET("/bulletin/list", bulletin.List)
		apiV1.DELETE("/bulletin/delete", bulletin.Delete)

		apiV1.POST("/dictType/create", dict.CreateDictType)
		apiV1.PUT("/dictType/update", dict.UpdateDictType)
		apiV1.GET("/dictType/get", dict.GetDictType)
		apiV1.GET("/dictType/list", dict.ListDictType)
		apiV1.DELETE("/dictType/delete", dict.DeleteDictType)
		apiV1.GET("/dictType/all", dict.AllDictType)

		apiV1.POST("/dictData/create", dict.CreateDictData)
		apiV1.PUT("/dictData/update", dict.UpdateDictData)
		apiV1.GET("/dictData/get", dict.GetDictData)
		apiV1.GET("/dictData/list", dict.ListDictData)
		apiV1.DELETE("/dictData/delete", dict.DeleteDictData)

		apiV1.POST("/dept/create", dept.Create)
		apiV1.PUT("/dept/update", dept.Update)
		apiV1.GET("/dept/get", dept.Get)
		apiV1.GET("/dept/list", dept.List)
		apiV1.DELETE("/dept/delete", dept.Delete)
	}
	//external
	externalRouter := g.Group("/external")
	{
		externalRouter.GET("/dictType/all", dict.AllDictType)
	}
	//multipartUpload
	multipartUploadRouter := g.Group("/multipartUpload")
	{
		multipartUploadRouter.POST("/*objectName", upload.InitOrCompleteMultipartUpload)
		multipartUploadRouter.PUT("/*objectName", upload.MultipartUploadPart)
		multipartUploadRouter.GET("/*objectName", upload.ListParts)
	}
	//The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/demo1", sd.DemoOne)
	}
	//public static
	publicRouter := g.Group("/public")
	{
		publicRouter.Static("", "public")
	}
	return g
}
