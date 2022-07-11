package spotify

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/lipefebn/spotify-crud/src/core/services/spotify"
)


func Artist(route fiber.Router) {
	route.Post("/", post)
}

// ShowAccount godoc
// @Summary      Insert album
// @Description  insert album in database
// @Tags         album
// @Accept       json
// @Produce      json
// @Param payload body repository.Album true "payload"
// @Success      201  {object}  handler.ResponseError
// @Failure      400  {object}  handler.ResponseError
// @Failure      404  {object}  handler.ResponseError
// @Failure      500  {object}  handler.ResponseError
// @Router       /album [post]
func post(c *fiber.Ctx) error {
	result := services.InitAlbumService(c.Body()).Select()
	return c.JSON(result)
}