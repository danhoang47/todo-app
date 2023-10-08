package usermodel

import (
	"errors"
	"fmt"
	"todo/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Username  string        `json:"username" gorm:"column:username;"`
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

func (u *User) GetUserId() int { return int(u.ID) }

func (u *User) GetUsername() string { return u.Username }

func (u *User) GetRole() string { return u.Role }

type UserLogin struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"-" gorm:"column:password;"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

func ErrUsernameExisted(username string) error {
	return fmt.Errorf("username %s has existed", username)
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomerError(
		errors.New("username or password is invalid"),
		"username or password is invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
