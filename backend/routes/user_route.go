package routes

import (
	"todo/component/appctx"
	"todo/middleware"
	"todo/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	{
		users.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.Profile(appCtx))
		users.POST("/register", ginuser.Register(appCtx))
		users.POST("/authenticate", ginuser.Login(appCtx))
	}
}
