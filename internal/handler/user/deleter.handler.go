package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"monopc-starter/common"
	userService "monopc-starter/internal/service/user"
	"monopc-starter/resource"
	"monopc-starter/utils"
)

type UserDeleterHandler struct {
	userDeleter  userService.UserDeleterUseCase
	cloudStorage utils.CloudStorage
}

func NewUserDeleterHandler(userDeleter userService.UserDeleterUseCase, cloudStorage utils.CloudStorage) *UserDeleterHandler {
	return &UserDeleterHandler{
		userDeleter:  userDeleter,
		cloudStorage: cloudStorage,
	}
}

func (udh *UserDeleterHandler) DeleteUser(c *gin.Context) {
	var request resource.DeleteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := uuid.Parse(request.ID)

	err = udh.userDeleter.DeleteUser(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (udh *UserDeleterHandler) DeleteRole(c *gin.Context) {
	var request resource.DeleteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := uuid.Parse(request.ID)

	err = udh.userDeleter.DeleteRole(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (udh *UserDeleterHandler) DeletePermission(c *gin.Context) {
	var request resource.DeleteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := uuid.Parse(request.ID)

	err = udh.userDeleter.DeletePermission(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}
