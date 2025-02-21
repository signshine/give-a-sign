package http

import (
	"fmt"

	"github.com/signshine/give-a-sign/app"

	"github.com/gofiber/fiber/v2"
)

func Run(app app.App) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAPI(api, app)

	return router.Listen(fmt.Sprintf(":%d", app.Config().Server.HttpPort))
}

func registerAPI(router fiber.Router, app app.App) {
	registerAuthAPI(router, app)
	registerLanguageAPI(router, app)
}

func registerAuthAPI(router fiber.Router, app app.App) {
	userSvcGetter := UserServiceGetter(app, app.Config().Server)
	router.Post("/sign-up", setTransaction(app.DB()), SignUp(userSvcGetter))
	router.Post("/sign-in", setTransaction(app.DB()), SignIn(userSvcGetter))
	secret := []byte(app.Config().Server.Secret)
	router.Get("/test-auth", newAuthMiddleware(secret), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "test-auth",
		})
	})
}

func registerLanguageAPI(router fiber.Router, app app.App) {
	languageSvcGetter := LanguageServiceGetter(app)
	secret := []byte(app.Config().Server.Secret)
	router.Use(newAuthMiddleware(secret))

	router.Post("/languages", setTransaction(app.DB()), CreateLanguage(languageSvcGetter))
	router.Get("/languages", setTransaction(app.DB()), GetListLanguage(languageSvcGetter))
	router.Get("/languages/filter", setTransaction(app.DB()), GetLanguage(languageSvcGetter))
	router.Delete("/languages", setTransaction(app.DB()), DeleteLanguage(languageSvcGetter))

	router.Post("/sign-languages", setTransaction(app.DB()), CreateSignLanguage(languageSvcGetter))
	router.Get("/sign-languages", setTransaction(app.DB()), GetListSignLanguage(languageSvcGetter))
	router.Get("/sign-languages/filter", setTransaction(app.DB()), GetSignLanguage(languageSvcGetter))
	router.Delete("/sign-languages", setTransaction(app.DB()), DeleteSignLanguage(languageSvcGetter))
}
