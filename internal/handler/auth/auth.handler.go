package auth

import (
	"github.com/gin-gonic/gin"
	"monopc-starter/common"
	"monopc-starter/internal/service/auth"
	"monopc-starter/resource"
	"net/http"
)

type AuthHandler struct {
	authService auth.AuthServiceUseCase
}

func NewAuthHandler(authService auth.AuthServiceUseCase) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var request resource.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	token, err := ah.authService.Login(c.Request.Context(), request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, common.ErrUnauthorized)
		return
	}

	c.JSON(200, resource.LoginResponse{
		Token: *token,
	})
}
