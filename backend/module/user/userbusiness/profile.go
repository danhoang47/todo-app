package userbusiness

import (
	"context"
	"todo/module/user/usermodel"
)

type ProfileStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfor ...string,
	) (*usermodel.User, error)
}

type profileBusiness struct {
	profileStorage ProfileStorage
}

func NewProfileBusiness(profileStorage ProfileStorage) *profileBusiness {
	return &profileBusiness{profileStorage}
}

func (biz *profileBusiness) Profile(ctx context.Context, userId string) (*usermodel.User, error) {

	return nil, nil
}
