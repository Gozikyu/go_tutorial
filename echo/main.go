package main

import (
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140"`
	UserName string `validate:"required,min=1,max=15"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Name: "John Doe", Email: "john@example.com"},
	{Name: "Jane Doe", Email: "jane@example.com"},
	// Add more users as needed
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	// Routes
	e.GET("/comments", func(c echo.Context) error {
		mutex.RLock()
		defer mutex.RUnlock()

		return c.JSON(http.StatusOK, comments)
	})

	e.POST("/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		validate := validator.New()
		if err := validate.Struct(comment); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"status": err.Error()})
		}

		mutex.Lock()
		comments = append(comments, comment)
		mutex.Unlock()

		return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
