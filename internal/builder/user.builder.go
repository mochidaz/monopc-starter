package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	user2 "monopc-starter/api/user"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/repository"
	userService "monopc-starter/internal/service/user"
	"monopc-starter/sdk/local"
)

func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	rr := repository.NewRoleRepository(db)
	urr := repository.NewUserRoleRepository(db)
	pr := repository.NewPermissionRepository(db)
	rpr := repository.NewRolePermissionRepository(db)

	uc := userService.NewUserCreator(cfg, ur, rr, urr, pr, rpr)
	uf := userService.NewUserFinder(ur, rr, pr, urr, rpr)
	ud := userService.NewUserDeleter(ur, urr, rr, rpr, pr)
	up := userService.NewUserUpdater(ur, rr, pr, urr, rpr)

	cloudStorage := local.NewLocalStorage(cfg.LocalStorage.BasePath)

	user2.UserCreatorHTTPHandler(cfg, router, uc, uf, cloudStorage)
	user2.UserFinderHTTPHandler(cfg, router, uc, uf, cloudStorage)
	user2.UserDeleterHTTPHandler(cfg, router, ud, uf, cloudStorage)
	user2.UserUpdaterHTTPHandler(cfg, router, up, uf, cloudStorage)
	user2.UserRegisterHTTPHandler(cfg, router, uc, uf, cloudStorage)
}
