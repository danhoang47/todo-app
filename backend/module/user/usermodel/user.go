package usermodel

import (
	"errors"
	"fmt"
	"todo/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"lastName" gorm:"column:last_name;"`
	FirstName string        `json:"firstName" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string { return "users" }

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"lastName" gorm:"column:last_name;"`
	FirstName string        `json:"firstName" gorm:"column:first_name;"`
	Role      string        `json:"_" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

func (u *User) GetUserId() int { return int(u.ID) }

func (u *User) GetEmail() string { return u.Email }

func (u *User) GetRole() string { return u.Role }

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

func ErrEmailExisted(email string) error {
	return fmt.Errorf("email %s has existed", email)
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomerError(
		errors.New("email or password is invalid"),
		"username or password is invalid",
		"ErrEmailOrPasswordInvalid",
	)
)
