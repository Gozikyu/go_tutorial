package main

import (
	"errors"
	"go_tutorial/echo/Infra"
	"go_tutorial/echo/Infra/ExternalServices"
	"go_tutorial/echo/UserController"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestEmail struct {
	Email string `validate:"required,min=1,max=140"`
}

func main() {
	db := ExternalServices.SetupTestDatabase()
	userRepository := Infra.NewUserRepository(db)
	companyRepository := Infra.NewCompanyRepository(db)
	userController := UserController.NewUserController(userRepository, companyRepository)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/**
	ユーザーを取得する
	*/
	e.GET("/user/:id", func(c echo.Context) error {
		userId := c.Param("id")
		intId, err := strconv.Atoi(userId)
		if err != nil {
			return errors.New("クエリパラメータのuserIdの変換に失敗しました")
		}
		user, err := userController.GetUser(intId)
		if err != nil {
			return errors.New("userの取得に失敗しました")
		}

		return c.JSON(http.StatusOK, user)
	})

	/**
	メールアドレスを変更する
	リクエスト例: curl -X PUT -H "Content-Type: application/json" -d '{"email": "new-email@example.com"}' http://localhost:8888/user/1
	*/
	e.PUT("/user/:id", func(c echo.Context) error {
		var email RequestEmail
		if err := c.Bind(&email); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
		}

		userId := c.Param("id")
		intId, err := strconv.Atoi(userId)
		if err != nil {
			return errors.New("クエリパラメータのuserIdの変換に失敗しました")
		}
		e := userController.ChangeEmail(intId, email.Email)
		if e != nil {
			return errors.New("情報の更新に失敗しました")
		}

		return c.JSON(http.StatusOK, "success")
	})

	/**
	メールアドレスを認証済み状態にする
	リクエスト例: curl -X PUT -H "Content-Type: application/json" http://localhost:8888/user/1/confirmEmail
	*/
	e.PUT("/user/:id/confirmEmail", func(c echo.Context) error {
		userId := c.Param("id")
		intId, err := strconv.Atoi(userId)
		if err != nil {
			return errors.New("クエリパラメータのuserIdの変換に失敗しました")
		}

		//本来は直接書き換えしてはいけない
		db.Users[intId]["isEmailConfirmed"] = "true"

		user, err := userController.GetUser(intId)
		if err != nil {
			return errors.New("userの取得に失敗しました")
		}
		return c.JSON(http.StatusOK, user)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
