package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/riyajath-ahamed/microfrontend-golang/internal/handlers"
)

func RegisterUserRoutes(g *echo.Group) {
	h := handlers.NewUserHandler()

	g.GET("/:id", h.GetUser)
	g.POST("/", h.CreateUser)
}
