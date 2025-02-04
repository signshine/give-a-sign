package http

import (
	"fmt"

	"github.com/signshine/give-a-sign/app"

	"github.com/gofiber/fiber/v2"
)

func Run(app app.App) error {
	router := fiber.New()

	api := router.Group("/api/v1", SetUserContext)

	registerAPI(api, app)

	return router.Listen(fmt.Sprintf(":%d", app.Config().Server.HttpPort))
}

func registerAPI(router fiber.Router, app app.App) {
	registerAuthAPI(router, app)
	// registerWordAPI(router, app)
}

func registerAuthAPI(router fiber.Router, app app.App) {
	userSvcGetter := UserServiceGetter(app, app.Config().Server)
	router.Post("/sign-up", SetTransaction(app.DB()), SignUp(userSvcGetter))
	router.Post("/sign-in", SetTransaction(app.DB()), SignIn(userSvcGetter))
}

// func registerWordAPI(router fiber.Router, app app.App) {
	// handler := NewWordHandler(service.NewWordService(app.WordService()))
	// wordSvcGetter := WordServiceGetter(app)
	// router.Post("/words", handler.AddWord)
	// router.Get("/words", handler.AllWords)
	// router.Post("words/:wordId/videos", handler.AddVideo)
	// router.Get("words/:wordId/videos", handler.AllVideos)
// }
