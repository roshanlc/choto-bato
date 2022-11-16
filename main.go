/*
Author: Roshan Lamichhane
Date: 16-11-2022

Choto-bato in Nepali means short(quick) route.
*/
package main

import (
	"net/http"

	"example.com/go-redis-fun/handler"
	"example.com/go-redis-fun/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Enable Global CORS support
	e.Use(middleware.CORS())

	// Handle the short url
	e.GET("/:url", func(c echo.Context) error {

		return handler.ReturnLongURL(c)
	})

	// Handle new short links creation request
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "Please use POST request to create new short urls."})
	})

	// Handle home page
	e.POST("/", func(c echo.Context) error {
		return handler.NewURLHandler(c)
	})

	// Start store facility
	store.InitializeStore()

	e.Logger.Fatal(e.Start(":9000"))
}
