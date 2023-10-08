package userbusiness

import (
	"context"
	"todo/common"
	"todo/module/user/usermodel"
)

type RegisterStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfor ...string,
	) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage,
		hasher,
	}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]any{"email": data.Email})

	if user != nil {
		// if user.Status == 0 {
		// 	return err user has been disabled
		// }

		return usermodel.ErrEmailExisted(user.Email)
	}

	salt := common.GetSalt(50)

	// Hashing
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
