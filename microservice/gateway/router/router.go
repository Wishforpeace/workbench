package router

import (
	"workbench-gateway/handler"
	"workbench-gateway/handler/sd"
	"workbench/pkg/errno"

	_ "workbench-gateway/docs"
	"workbench-gateway/handler/user"
	"workbench-gateway/router/middleware"
	"workbench/pkg/constvar"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		handler.SendError(c, errno.ErrIncorrectAPIRoute, nil, "", "")
	})

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 权限要求，普通用户/管理员/超管
	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	// adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	// superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// project 权限
	// projectCheck := middleware.ProjectMiddleware()

	// auth 模块
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/login", user.Login)
		authRouter.POST("/signup", user.Register)
	}

	// user 模块
	userRouter := g.Group("api/v1/user")
	userRouter.Use(normalRequired)
	{
		// userRouter.GET("/infos", user.GetInfo)
		userRouter.GET("/profile/:id", user.GetProfile)
		userRouter.GET("/myprofile", user.GetMyProfile)
		userRouter.GET("/list/:group_id/:team_id", user.List)
		userRouter.PUT("", user.UpdateInfo)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
