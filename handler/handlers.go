package handler

import (
	"net/http"

	"example.com/go-redis-fun/shortener"
	"example.com/go-redis-fun/store"
	"github.com/labstack/echo/v4"
)

// Struct to parse request body
type URLCreationRequest struct {
	LongURL string `json:"long_url",binding:"required`
}

const host = "http://localhost:9000/"

// Echo handler to handle POST requests to create new short urls
func NewURLHandler(c echo.Context) error {
	body := new(URLCreationRequest)

	if err := c.Bind(body); err != nil {
		return err
	}

	// Generate a short url
	shortURL, err := shortener.GenerateShortURL([]byte(body.LongURL))

	if err != nil {
		return err
	}

	// Save in redis
	store.SaveURL(shortURL, body.LongURL)

	return c.JSON(http.StatusOK, echo.Map{"short_url": host + shortURL,
		"long_url": body.LongURL})
}

// Echo handler to handle get requests to  retrieve long urls
func ReturnLongURL(c echo.Context) error {

	shortURL := c.Param("url")

	// Fetch from redis
	longURL, err := store.RetrieveURL(shortURL)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"short_url": host + shortURL,
		"long_url": longURL})
}
