package main

import (
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

	e.GET("/company", func(c echo.Context) error {

	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
