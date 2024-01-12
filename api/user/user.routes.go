package user

import (
	"github.com/gin-gonic/gin"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/app/middleware"
	"monopc-starter/internal/handler/auth"
	"monopc-starter/internal/handler/user"
	auth2 "monopc-starter/internal/service/auth"
	userService "monopc-starter/internal/service/user"
	"monopc-starter/utils"
)

func UserCreatorHTTPHandler(cfg config.Config, router *gin.Engine, uc userService.UserCreatorUseCase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := user.NewUserCreatorHandler(uc, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))
	{
		api.POST("/cms/user", hndlr.CreateUser)
		api.POST("/cms/admin", hndlr.CreateAdmin)
		api.POST("/cms/role", hndlr.CreateRole)
		api.POST("/cms/permission", hndlr.CreatePermission)
	}
}

func UserRegisterHTTPHandler(cfg config.Config, router *gin.Engine, uc userService.UserCreatorUseCase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := user.NewUserCreatorHandler(uc, cloudStorage)

	api := router.Group("/api")

	api.POST("/register", hndlr.CreateUser)
}

func UserFinderHTTPHandler(cfg config.Config, router *gin.Engine, uc userService.UserCreatorUseCase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := user.NewUserFinderHandler(uf, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))
	{
		api.GET("/cms/user", hndlr.FindUsers)
		api.GET("/cms/role", hndlr.FindRoles)
		api.GET("/cms/permission", hndlr.FindPermissions)
		api.GET("/cms/user/:id", hndlr.FindUserById)
		api.GET("/cms/role/:id", hndlr.FindRoleById)
		api.GET("/cms/permission/:id", hndlr.FindPermissionById)
	}
}

func UserDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ud userService.UserDeleterUseCase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := user.NewUserDeleterHandler(ud, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))
	{
		api.DELETE("/cms/user", hndlr.DeleteUser)
		api.DELETE("/cms/role", hndlr.DeleteRole)
		api.DELETE("/cms/permission", hndlr.DeletePermission)
	}
}

func UserUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uu userService.UserUpdaterUseCase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := user.NewUserUpdaterHandler(uu, cloudStorage)

	api := router.Group("/api")

	{
		api.PUT("/user/update", hndlr.UpdateUser)
		api.PUT("/user/update/password", hndlr.UpdatePassword)
	}

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))
	{
		api.PUT("/cms/user", hndlr.UpdateUser)
		api.PUT("/cms/role", hndlr.UpdateRole)
		api.PUT("/cms/permission", hndlr.UpdatePermission)
	}
}

func UserAuthHTTPHandler(cfg config.Config, router *gin.Engine, as auth2.AuthServiceUseCase) {
	hndlr := auth.NewAuthHandler(as)

	api := router.Group("/api")
	{
		api.POST("/login", hndlr.Login)
	}
}
