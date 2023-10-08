package middleware

import (
	"todo/common"
	"todo/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, allowRole := range allowRoles {
			if user.GetRole() == allowRole {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNoPermission(nil))
		}

		c.Next()
	}
}
