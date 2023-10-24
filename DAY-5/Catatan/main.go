package main

import (
	"github.com/labstack/echo/v4"
)

// type User struct {
// 	Id    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

//	func GetUsers(c echo.Context) error {
//		var user User = User{1, "Sayid", "Sayid@gmail.com"}
//		return c.JSON(200, user)
//	}
//
//	func GetUsersController(c echo.Context) error {
//		id, _ := strconv.Atoi(c.Param("id"))
//		user := User{Id: id, Name: "Sayidaaaa", Email: "Sayisdd@gmail.com"}
//		return c.JSON(200, map[string]interface{}{"user": user})
//	}
//
//	func UserSearchController(c echo.Context) error {
//		match := c.QueryParam("match")
//		return c.JSON(200, map[string]interface{}{
//			"match":  match,
//			"result": []string{"adis", "aasn", "aes"},
//		})
//	}
type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CreateBindUser(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	return c.JSON(200, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

// func CreateFormUser(c echo.Context) error {
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")

//		var user User = User{Name: name, Email: email}
//		return c.JSON(200, map[string]interface{}{
//			"message": "success create user",
//			"user":    user,
//		})
//	}
func main() {
	e := echo.New()
	e.POST("/users", CreateBindUser)
	e.Start(":8000")
	// e := echo.New()
	// e.GET("/users", GetUsers)
	// e.GET("/users:id", GetUsersController)
	// e.GET("/users", UserSearchController)
	// e.Start(":8000")
}
