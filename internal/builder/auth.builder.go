package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monopc-starter/api/user"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/repository"
	auth2 "monopc-starter/internal/service/auth"
)

func BuildAuthHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ar := repository.NewAuthRepository(db)
	ur := repository.NewUserRepository(db)

	as := auth2.NewAuthService(cfg, ur, ar)

	user.UserAuthHTTPHandler(cfg, router, as)
}
