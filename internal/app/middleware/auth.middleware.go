package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"monopc-starter/internal/app/config"
	"monopc-starter/utils"
	"net/http"
	"strings"
)

var UserID uuid.UUID

func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, utils.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ErrorApiResponse(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		UserID = claims.Subject

		c.Next()
	}
}

func Admin(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, utils.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ErrorApiResponse(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		if claims.Issuer != cfg.JWTConfig.IssuerCMS {
			c.JSON(http.StatusUnauthorized, utils.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		UserID = claims.Subject

		c.Next()
	}
}
