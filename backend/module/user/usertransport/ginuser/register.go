package ginuser

import (
	"log"
	"net/http"
	"todo/common"
	"todo/component/appctx"
	"todo/component/hasher"
	"todo/module/user/userbusiness"
	"todo/module/user/usermodel"
	"todo/module/user/userstorage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.User

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		log.Println(data)

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(data.Username, err))
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
