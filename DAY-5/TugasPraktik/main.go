package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts(c echo.Context) error {
	// Kirim permintaan GET ke API JSONPlaceholder
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Baca data dari respons
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Kirim data ke klien
	return c.JSONBlob(http.StatusOK, body)
}

func GetPostByID(c echo.Context) error {
	id := c.Param("id")
	url := "https://jsonplaceholder.typicode.com/posts/" + id

	// Kirim permintaan GET dengan ID tertentu ke API JSONPlaceholder
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Baca data dari respons
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Kirim data ke klien
	return c.JSONBlob(http.StatusOK, body)
}

func CreatePost(c echo.Context) error {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}

	// Kirim data ke API JSONPlaceholder menggunakan permintaan POST
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Proses respons jika diperlukan
	// ...

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post created successfully",
		"post":    post,
	})
}

func DeletePost(c echo.Context) error {
	id := c.Param("id")
	url := "https://jsonplaceholder.typicode.com/posts/" + id

	// Buat permintaan HTTP DELETE ke API JSONPlaceholder
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Proses respons jika diperlukan
	// ...

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post deleted successfully",
	})
}

func main() {
	e := echo.New()

	// Daftarkan rute-rute yang diperlukan
	e.GET("/posts", GetPosts)
	e.GET("/posts/:id", GetPostByID)
	e.POST("/posts", CreatePost)
	e.DELETE("/posts/:id", DeletePost)

	// Mulai server Echo pada port 8000
	e.Start(":8000")

	// Sisanya dari kode Anda...
	// Pastikan Anda memiliki struktur data yang sesuai dan menyesuaikan kode
	// di dalam kontroler sesuai kebutuhan aplikasi Anda.
}
