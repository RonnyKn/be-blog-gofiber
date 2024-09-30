package controller

import (
	"log"

	"github.com/RonnyKn/be-blog-gofiber/db"
	"github.com/RonnyKn/be-blog-gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog List",
	}

	db := db.DBConn
	var blogs []models.Blog
	db.Find(&blogs)

	context["data"] = blogs

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{}

	data := new(models.Blog)
	if err := c.BodyParser(&data); 
	 err != nil {
		log.Println("Error in parsing request", err)
		context["message"] = err
	}

  	result := db.DBConn.Create(&data)

	// if data.BlogID == 0 {
	// 	log.Println("Error in creating blog", result.Error)
	// 	context["message"] = result.Error
	// }

  	if result.Error != nil {
		log.Println("Error in creating blog", result.Error)
		context["message"] = result.Error
	}
	context["data"] = data
	context["message"] = "Blog Created"
	c.Status(200)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog UPDATED",
	}

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog Deleted",
	}

	c.Status(200)
	return c.JSON(context)
}