package api

import (
	auth "dapoint-api/api/v1/auth"
	contentV1 "dapoint-api/api/v1/content"
	"dapoint-api/api/v1/user"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller
	UserController      *user.Controller
	AuthController      *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	// contentV1.Use(middleware.JWTMiddleware())
	contentV1.GET("", controller.ContentV1Controller.GetAll)
	contentV1.POST("", controller.ContentV1Controller.Create)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)

	user := e.Group("/users")
	user.GET("", controller.UserController.GetAll)
	user.GET("/:id", controller.UserController.GetByID)
	user.POST("", controller.UserController.Create)
}
