package ginuser

import (
	"net/http"
	"todo/common"
	"todo/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
