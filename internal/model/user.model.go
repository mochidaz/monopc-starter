package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

const (
	userTableName = "users"
)

type User struct {
	ID                  uuid.UUID    `json:"id"`
	Name                string       `json:"name"`
	Email               string       `json:"email"`
	Password            string       `json:"password"`
	PhoneNumber         string       `json:"phone_number"`
	Photo               string       `json:"photo"`
	DOB                 sql.NullTime `json:"dob"`
	ForgotPasswordToken string       `json:"forgot_password_token"`
	UserRole            *UserRole    `json:"user_role"`
	Auditable
}

func (*User) TableName() string {
	return userTableName
}

func NewUser(
	id uuid.UUID,
	name string,
	email string,
	password string,
	phoneNumber string,
	photo string,
	dob sql.NullTime,
	createdBy string,
) *User {
	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Photo:       photo,
		DOB:         dob,
		Auditable:   NewAuditable(createdBy),
	}
}

func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":         model.Name,
			"email":        model.Email,
			"password":     model.Password,
			"phone_number": model.PhoneNumber,
			"photo":        model.Photo,
			"dob":          model.DOB,
			"updated_at":   model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Email != from.Email {
		mapped["email"] = from.Email
	}

	if model.Password != from.Password {
		mapped["password"] = from.Password
	}

	if model.PhoneNumber != from.PhoneNumber {
		mapped["phone_number"] = from.PhoneNumber
	}

	if model.Photo != from.Photo {
		mapped["photo"] = from.Photo
	}

	if model.DOB != from.DOB {
		mapped["dob"] = from.DOB
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
