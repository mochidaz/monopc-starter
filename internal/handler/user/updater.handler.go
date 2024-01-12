package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"monopc-starter/common"
	"monopc-starter/internal/model"
	userService "monopc-starter/internal/service/user"
	"monopc-starter/resource"
	"monopc-starter/utils"
	"strconv"
)

type UserUpdaterHandler struct {
	userUpdater  userService.UserUpdaterUseCase
	cloudStorage utils.CloudStorage
}

func NewUserUpdaterHandler(userUpdater userService.UserUpdaterUseCase, cloudStorage utils.CloudStorage) *UserUpdaterHandler {
	return &UserUpdaterHandler{
		userUpdater:  userUpdater,
		cloudStorage: cloudStorage,
	}
}

func (uh *UserUpdaterHandler) UpdateUser(c *gin.Context) {
	var request resource.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	imagePath, err := uh.cloudStorage.UploadFile(request.Photo, "users/user/profile")

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	dob := utils.ParseDate(request.DOB)

	id, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	newData := model.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Photo:       imagePath,
		DOB:         sql.NullTime{Time: dob, Valid: true},
	}

	err = uh.userUpdater.UpdateUser(c.Request.Context(), id, &newData)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (uh *UserUpdaterHandler) UpdateRole(c *gin.Context) {
	var request resource.UpdateRoleRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := strconv.Atoi(request.ID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	newData := model.Role{
		Name: request.Name,
	}

	err = uh.userUpdater.UpdateRole(c.Request.Context(), id, &newData)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (uh *UserUpdaterHandler) UpdatePermission(c *gin.Context) {
	var request resource.UpdatePermissionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := strconv.Atoi(request.ID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	newData := model.Permission{
		Name:  request.Name,
		Label: request.Label,
	}

	err = uh.userUpdater.UpdatePermission(c.Request.Context(), id, &newData)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, nil)
}

func (uh *UserUpdaterHandler) UpdatePassword(context *gin.Context) {
	var request resource.UpdatePasswordRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, common.ErrBadRequest)
		return
	}

	id, err := uuid.Parse(request.ID)

	if err != nil {
		context.JSON(400, common.ErrBadRequest)
		return
	}

	err = uh.userUpdater.UpdatePassword(context.Request.Context(), id, request.OldPassword, request.NewPassword)

	if err != nil {
		context.JSON(400, common.ErrBadRequest)
		return
	}

	context.JSON(200, nil)
}
