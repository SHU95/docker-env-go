package infrastructure

import (
	"github.com/SHU95/docker-env-go/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewMySqlDb())

	//ログを出力
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.GET("/users/:id", func(context echo.Context) error { return userController.GetUser(context) })

	e.POST("/user/create", func(context echo.Context) error { return userController.CreateUser(context) })

	e.PUT("/users/:id", func(context echo.Context) error { return userController.UpdateUser(context) })

	e.Logger.Fatal(e.Start(":8080"))
}
