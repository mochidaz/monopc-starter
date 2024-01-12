package user

import (
	"github.com/gin-gonic/gin"
	"monopc-starter/common"
	userService "monopc-starter/internal/service/user"
	"monopc-starter/resource"
	"monopc-starter/utils"
	"strconv"
)

type UserCreatorHandler struct {
	userCreator  userService.UserCreatorUseCase
	cloudStorage utils.CloudStorage
}

func NewUserCreatorHandler(userCreator userService.UserCreatorUseCase, cloudStorage utils.CloudStorage) *UserCreatorHandler {
	return &UserCreatorHandler{
		userCreator:  userCreator,
		cloudStorage: cloudStorage,
	}
}

func (uch *UserCreatorHandler) CreateUser(c *gin.Context) {
	var request resource.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	imagePath, err := uch.cloudStorage.UploadFile(request.Photo, "users/user/profile")

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	dob := utils.ParseDate(request.DOB)

	user, err := uch.userCreator.CreateUser(c.Request.Context(), request.Name, request.Email, request.Password, request.PhoneNumber, imagePath, dob)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, user)
}

func (uch *UserCreatorHandler) CreateAdmin(c *gin.Context) {
	var request resource.CreateAdminRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	imagePath, err := uch.cloudStorage.UploadFile(request.Photo, "users/user/profile")

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	dob := utils.ParseDate(request.DOB)

	roleID, err := strconv.Atoi(request.RoleID)

	user, err := uch.userCreator.CreateAdmin(c.Request.Context(), request.Name, request.Email, request.Password, request.PhoneNumber, imagePath, dob, roleID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, user)
}

func (uch *UserCreatorHandler) CreateRole(c *gin.Context) {
	var request resource.CreateRoleRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	permissions := make([]int, 0)

	for _, permissionID := range request.PermissionIDs {
		permission, err := strconv.Atoi(permissionID)

		if err != nil {
			c.JSON(400, common.ErrBadRequest)
			return
		}

		permissions = append(permissions, permission)
	}

	role, err := uch.userCreator.CreateRole(c.Request.Context(), request.Name, permissions, "system")

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, role)
}

func (uch *UserCreatorHandler) CreatePermission(c *gin.Context) {
	var request resource.CreatePermissionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	permission, err := uch.userCreator.CreatePermission(c.Request.Context(), request.Name, request.Label)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, permission)
}
