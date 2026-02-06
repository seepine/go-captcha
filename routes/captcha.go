package routes

import (
	"fmt"
	"go-captcha/common"
	"go-captcha/g"
	"go-captcha/utils"
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v5"
)

var ApiKeys = func() []string {
	keys := utils.Getenv("API_KEYS", "")
	if keys == "" {
		return []string{}
	}
	return strings.Split(keys, ",")
}()

func CaptchaGen(c *echo.Context) error {
	apiKey := c.Request().Header.Get("ApiKey")

	// 验证 API 密钥中是否包含 req.ApiKey，空表示不鉴权
	if len(ApiKeys) > 0 {
		found := slices.Contains(ApiKeys, apiKey)
		if !found {
			return g.FailWithCode(c, http.StatusUnauthorized, "无效的 API 密钥")
		}
	}

	slideCaptcha := common.CaptchaSlide()

	captData, err := slideCaptcha.Generate()
	if err != nil {
		fmt.Println(err)
		return g.Fail(c, "生成验证码失败")
	}

	fmt.Println("123")
	dotData := captData.GetData()
	fmt.Println(dotData)
	if dotData == nil {
		return g.Fail(c, "生成验证码失败[2]")
	}

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		fmt.Println(err)
		return g.Fail(c, "生成验证码图片失败")
	}
	tBase64, err = captData.GetTileImage().ToBase64()
	if err != nil {
		fmt.Println(err)
		return g.Fail(c, "生成验证码图片[2]失败")
	}

	return g.Ok(c, g.M{
		"slideData": g.M{
			"thumbX":      dotData.DX,
			"thumbY":      dotData.DY,
			"thumbWidth":  dotData.Width,
			"thumbHeight": dotData.Height,
			"image":       mBase64,
			"thumb":       tBase64,
		},
		"slideVerifyData": g.M{
			"x": dotData.X,
			"y": dotData.Y,
		},
	})

}
