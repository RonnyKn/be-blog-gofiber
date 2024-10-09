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

  	if result.Error != nil {
		log.Println("Error in creating blog", result.Error)
		context["message"] = result.Error
	}
	context["data"] = data
	context["message"] = "Blog Created"
	c.Status(200)
	return c.JSON(context)
}

func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}

	id := c.Params("id")
	var blog models.Blog
	
	db.DBConn.First(&blog, id)

	if blog.BlogID == 0 {
		log.Println("Error, Blog " + id + " Not Found")
		context["message"] = "Blog " + id + " Not Found"
		c.Status(404)
		return c.JSON(context)
	}
	c.Status(200)
	context["status Text"] = "OK"
	context["message"] = "Blog Details"
	context["data"] = blog
	return c.JSON(context)
}
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog UPDATED",
	}

	id := c.Params("id")
	
	db := db.DBConn
	var blog models.Blog
	data := db.First(&blog, id)

	if blog.BlogID == 0 {
		log.Println("Blog Not Found")
		context["message"] = "Blog "+id+" Not Found"
		c.Status(404)
		return c.JSON(context)
	}

	if data.Error != nil {
		log.Println("Error in updating blog", data.Error)
		context["message"] = data.Error
	}

	if err := c.BodyParser(&blog); err != nil {
		log.Println("Error in parsing request", err)
	}
	result := db.Save(&blog)

	if result.Error != nil {
		log.Println("Error in updating blog", result.Error)
		context["message"] = result.Error}
	c.Status(200)
	context["data"] = blog
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog Deleted",
	}

	id := c.Params("id")
	db := db.DBConn
	var blog models.Blog
	data := db.First(&blog, id)

	if blog.BlogID == 0{
		log.Println("Blog "+id+" Not Found")
		context["message"] = "Blog "+id+" Not Found"
		c.Status(404)
		return c.JSON(context)	
	}
	if data.Error != nil{
		log.Println("Error in deleting blog", data.Error)
		context["message"] = data.Error
	}	

	result := db.Delete(&blog, id)
	if result.Error != nil{
		log.Println("Error in deleting blog", result.Error)
		context["message"] = result.Error
	} 
		context["data"] = blog
		context["message"] = "Blog "+id+" Deleted"
		c.Status(200)
		return c.JSON(context)
	}
