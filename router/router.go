package router

import (
	"github.com/RonnyKn/be-blog-gofiber/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	//get route
	app.Get("/", controller.BlogList)

	//post route
	app.Post("/", controller.BlogCreate)

	//put route
	app.Put("/:id", controller.BlogUpdate)

	//delete route	
	app.Delete("/:id", controller.BlogDelete)
}