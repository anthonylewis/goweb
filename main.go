package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func usersIndex(c *gin.Context) {
	var users []User

	db.Find(&users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	count := 0
	db.Table("users").Count(&count)

	if count == 0 {
		db.Create(&User{Name: "Alice"})
		db.Create(&User{Name: "Bob"})
		db.Create(&User{Name: "Carol"})
	}

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", usersIndex)

	router.Run()
}
