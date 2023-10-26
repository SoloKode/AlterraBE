package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	InitDB()
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetDetailUsersController)
	e.POST("/users", AddUserController)
	e.Start(":8000")
}

var DB *gorm.DB

func InitDB() {
	// Ganti informasi koneksi sesuai dengan database Anda
	dsn := "root:@tcp(127.0.0.1:3306)/prakerja13?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&User{})
}
func AddUserController(c echo.Context) error {
	var user User
	c.Bind(&user)
	result := DB.Create(&user)
	if result.Error != nil {
		return c.JSON(500, BaseResponse{
			Status:  false,
			Message: "Failed create in database",
			Data:    nil,
		})
	}
	return c.JSON(200, BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    user,
	})
}

func GetUsersController(c echo.Context) error {
	// Ganti dengan kode yang mengambil data pengguna dari database
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(500, BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	return c.JSON(200, BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    users,
	})
}

func GetDetailUsersController(c echo.Context) error {
	id := c.Param("id")
	// Ganti dengan kode yang mengambil data pengguna berdasarkan ID dari database
	var user User
	DB.First(&user, id)
	return c.JSON(200, user)
}
