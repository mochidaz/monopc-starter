package resource

import "mime/multipart"

type CreateUserRequest struct {
	Email       string                `json:"email" form:"email" binding:"required"`
	Password    string                `json:"password" form:"password" binding:"required"`
	Name        string                `json:"name" form:"name" binding:"required"`
	Photo       *multipart.FileHeader `json:"photo" form:"photo"`
	DOB         string                `json:"dob" form:"dob"`
	PhoneNumber string                `json:"phone_number" form:"phone_number"`
}

type CreateUserResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	DOB         string `json:"dob"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type CreateAdminRequest struct {
	Email       string                `json:"email" binding:"required"`
	Password    string                `json:"password" binding:"required"`
	Name        string                `json:"name" binding:"required"`
	Photo       *multipart.FileHeader `json:"photo"`
	DOB         string                `json:"dob"`
	PhoneNumber string                `json:"phone_number"`
	RoleID      string                `json:"role_id"`
}

type CreateRoleRequest struct {
	Name          string   `json:"name" binding:"required"`
	PermissionIDs []string `json:"permission_ids"`
}

type CreatePermissionRequest struct {
	Name  string `json:"name" binding:"required"`
	Label string `json:"label" binding:"required"`
}

type CreatePermissionResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type CreateRoleResponse struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	PermissionIDs []string `json:"permission_ids"`
}

type UpdateUserRequest struct {
	ID          string                `json:"id" binding:"required"`
	Email       string                `json:"email" binding:"required"`
	Password    string                `json:"password" binding:"required"`
	Name        string                `json:"name" binding:"required"`
	Photo       *multipart.FileHeader `json:"photo"`
	DOB         string                `json:"dob"`
	PhoneNumber string                `json:"phone_number"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

type UpdateRoleRequest struct {
	ID            string   `json:"id" binding:"required"`
	Name          string   `json:"name" binding:"required"`
	PermissionIDs []string `json:"permission_ids"`
}

type UpdatePermissionRequest struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Label string `json:"label" binding:"required"`
}

type FindByIDRequest struct {
	ID string `json:"id"`
}

type UpdatePasswordRequest struct {
	ID          string `json:"id" binding:"required"`
	NewPassword string `json:"password" binding:"required"`
	OldPassword string `json:"old_pass" binding:"required"`
}
