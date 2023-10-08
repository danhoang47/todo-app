package ginuser

import (
	"net/http"
	"todo/common"
	"todo/component/appctx"
	"todo/component/hasher"
	"todo/component/tokenprovider/jwt"
	"todo/module/user/userbusiness"
	"todo/module/user/usermodel"
	"todo/module/user/userstorage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()

		loginBusiness := userbusiness.NewLoginBusiness(store, tokenProvider, md5, 60*60)

		account, err := loginBusiness.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
