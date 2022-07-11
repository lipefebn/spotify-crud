package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipefebn/spotify-crud/src/core/handler/spotify"
)

func Init() *fiber.App {
	app := fiber.New()
	spotify.Artist(app.Group("album"))
	return app
}
