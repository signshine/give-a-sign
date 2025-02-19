package http

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/signshine/give-a-sign/pkg/context"
	appJWT "github.com/signshine/give-a-sign/pkg/jwt"
	"github.com/signshine/give-a-sign/pkg/logger"
	"gorm.io/gorm"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext(), context.WithLogger(logger.NewLogger())))
	return c.Next()
}

func setTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tx := db.Begin()

		context.SetDB(c.UserContext(), tx, true)

		err := c.Next()

		if c.Response().StatusCode() >= 300 {
			return context.Rollback(c.UserContext())
		}

		if err := context.CommitOrRollback(c.UserContext(), true); err != nil {
			return err
		}

		return err
	}

}

func newAuthMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &appJWT.UserClaims{},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SuccessHandler: func(c *fiber.Ctx) error {
			claims := userClaims(c)
			if claims == nil {
				return fiber.ErrUnauthorized
			}

			logger := context.GetLogger(c.UserContext())
			context.SetLogger(c.UserContext(), logger.With("user_id", claims.UserID))

			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
	})
}

func userClaims(ctx *fiber.Ctx) *appJWT.UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*jwt.Token).Claims.(*appJWT.UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}
