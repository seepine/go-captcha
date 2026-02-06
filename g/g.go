package g

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type M map[string]any

// Ok 成功 json
func Ok(c *echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

// Fail 失败 json
func Fail(c *echo.Context, msg string) error {
	return c.JSON(http.StatusInternalServerError, M{"message": msg})
}

// Fail 失败 json
func FailWithCode(c *echo.Context, code int, msg string) error {
	return c.JSON(code, M{"message": msg})
}
