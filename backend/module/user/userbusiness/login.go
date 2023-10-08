package userbusiness

import (
	"context"
	"log"
	"todo/common"
	tokenprovider "todo/component/tokenprovider"
	"todo/module/user/usermodel"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfor ...string,
	) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]any{"username": data.Username})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	log.Println(data.Password)
	log.Println(user.Salt)
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	log.Println(passHashed)

	if passHashed != user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
