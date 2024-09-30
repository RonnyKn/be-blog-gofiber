package db

import (
	"log"
	"os"

	"github.com/RonnyKn/be-blog-gofiber/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := db_user+":"+db_pass+"@tcp(127.0.0.1:3306)/"+db_name+"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	  if err != nil {
    panic("failed to connect database")
  }
  log.Println("Connection Opened to Database")

  db.AutoMigrate(new(models.Blog))

  DBConn = db
}