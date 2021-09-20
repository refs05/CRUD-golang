package routes

import (
	"crudgoeg/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()

	echoInit.GET("/user", controllers.GetUserController)
	echoInit.POST("/user", controllers.PostUserController)
	echoInit.GET("/user/:id", controllers.GetbyIdUserController)
	echoInit.PUT("/user/:id", controllers.UpdateUserController)
	echoInit.DELETE("/user/:id", controllers.DeleteUserController)

	return echoInit
}

