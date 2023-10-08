package middleware

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"todo/common"
	"todo/component/appctx"
	"todo/component/tokenprovider/jwt"
	"todo/module/user/userstorage"

	"github.com/gin-gonic/gin"
)

var (
	ErrWrongAuthHeader = errors.New("wrong auth header")
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader
	}

	return parts[1], nil
}

func RequiredAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		storage := userstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := storage.FindUser(c.Request.Context(), map[string]any{"id": payload.UserId})

		log.Println(user)
		if err != nil {
			panic(common.ErrEntityNotFound(fmt.Sprint(payload.UserId), err))
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user not have permission")))
		}

		// Mask set fakeId
		user.Mask(false)

		// Set user to ginContext
		c.Set(common.CurrentUser, user)

		c.Next()
	}
}
