package main

import (
	//"github.com/lipefebn/spotify-crud/src/cli"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/lipefebn/spotify-crud/docs"
	"github.com/lipefebn/spotify-crud/src/core/handler"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := handler.Init()
	app.Get("/swagger/*", swagger.HandlerDefault)
	
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	app.Listen(":8080")
	//cli.Init()
}
