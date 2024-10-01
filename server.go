package main

import (
	"log"

	"github.com/RonnyKn/be-blog-gofiber/db"
	"github.com/RonnyKn/be-blog-gofiber/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")}

	//db connection
	db.ConnectDB()}

func main() {

	// Get DB connection
	sqlDB, err := db.DBConn.DB()
	if err != nil {
		panic(err)}

	defer sqlDB.Close()

	// Create server
	app := fiber.New()

	// Middlewares logger
	app.Use(logger.New())

	// CORS
	app.Use(cors.New())

	// Routes
	router.SetupRouter(app)

	// Listen and Server in :4000
	app.Listen(":4000")

	
}