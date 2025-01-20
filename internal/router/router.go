package router

import (
	_ "github.com/cbglabs/nexusgate-pbmapi/api/swagger"
	auth "github.com/cbglabs/nexusgate-pbmapi/internal/middleware"
	guiltysparkAPI "github.com/cbglabs/nexusgate-pbmapi/internal/router/guiltyspark_api"
	portalAPI "github.com/cbglabs/nexusgate-pbmapi/internal/router/portal_api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// MountRoutes sets up all the routers for the application
func MountRoutes() *gin.Engine {
	// Create a new Gin router, attach logger and recovery middleware
	router := gin.Default()

	// Add a home route
	router.GET("", func(context *gin.Context) {
		context.JSON(200, "Welcome home")
	})

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")

	// Portal API routes
	portalGroup := v1.Group("/portal-api")
	portalAPI.NewHealthRouter(portalGroup)
	portalAPI.NewAuthRouter(portalGroup)
	portalAPI.NewPublicRouter(portalGroup)
	portalAPI.NewUserRouter(portalGroup, auth.AuthSimpleMiddleware())

	// Cervey API routes
	cerveyGroup := v1.Group("/cervey-api")
	portalAPI.NewHealthRouter(cerveyGroup)
	// Add other Cervey-specific routes here

	// Guiltyspark API routes
	guiltysparkGroup := v1.Group("/guiltyspark-api")
	portalAPI.NewHealthRouter(guiltysparkGroup)
	guiltysparkAPI.NewMemberRouter(guiltysparkGroup, auth.AuthSimpleMiddleware())
	// Add other Guiltyspark-specific routes here

	return router
}
