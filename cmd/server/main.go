package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	authmw "github.com/riyajath-ahamed/microfrontend-golang/internal/middleware"
	"github.com/riyajath-ahamed/microfrontend-golang/internal/routes"
	// "github.com/riyajath-ahamed/microfrontend-golang/internal/handlers"
)

func main() {
	fmt.Println(" Server is running ")

	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	apiToken := os.Getenv("MICROFRONTEND_API_TOKEN")

	protected := e.Group("/api")
	protected.Use(authmw.TokenAuthMiddleware(apiToken))

	userGroup := protected.Group("/user")
	routes.RegisterUserRoutes(userGroup)

	//add user data and featureFlags

	e.Logger.Fatal(e.Start(":8080"))
}
