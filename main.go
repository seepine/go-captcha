package main

import (
	"embed"
	"go-captcha/common"
	"go-captcha/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

//go:embed assets
var webAssets embed.FS

func main() {
	godotenv.Load()
	common.InitLog()
	common.InitSnowflareId()
	common.InitCaptcha()

	e := echo.New()
	// middleware
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      false,
		Root:       "assets", // because files are located in `assets` directory in `webAssets` fs
		Filesystem: webAssets,
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 10 << 10, // 10 KB
	}))

	// register routes
	RegisterRoutes(e)

	log.Error().Err(e.Start(":" + utils.Getenv("PORT", "8080"))).Send()
}
