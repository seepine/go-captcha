package main

import (
	"go-captcha/g"
	"go-captcha/routes"
	"go-captcha/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

// health check
var startAt = time.Now().Format(time.DateTime)

func healthCheck(c *echo.Context) error {
	return c.JSON(http.StatusOK, g.M{
		"status":  "ok",
		"startAt": startAt,
	})
}

func RegisterRoutes(app *echo.Echo) {
	app.GET("/health", healthCheck)

	api := app.Group(utils.Getenv("BASE_PATH", "/api"))
	// captcha
	{
		captcha := api.Group("/captcha")
		captcha.Match([]string{"GET", "POST"}, "/gen", routes.CaptchaGen)
	}
}
