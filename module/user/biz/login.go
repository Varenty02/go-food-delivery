package userbiz

import (
	"context"
	"g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/component/tokenprovider"
	usermodel "g05-fooddelivery/module/user/model"
	"log"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {

	return &loginBusiness{storeUser: storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}
func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	log.Println(user)
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	passHashed := business.hasher.Hash(data.Password + user.Salt)
	log.Println(user.Password, passHashed, data.Password)
	if user.Password != passHashed {

		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{UserId: user.Id, Role: user.Role}
	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
