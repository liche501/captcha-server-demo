package main

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	DefaultLen = 4
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	e.GET("/key", APICaptchaKey)
	e.GET("/image/:key", APICaptchaImage)
	e.POST("/verify", APICaptchaVerify)
	e.POST("/reload/:key", APICaptchaReload)

	e.Start(":8888")
}

//获取key
func APICaptchaKey(c echo.Context) error {
	d := struct {
		CaptchaKey string
	}{
		// captcha.New(),
		captcha.NewLen(DefaultLen),
	}
	// fmt.Println("CaptchaId-->", d)

	return c.JSON(http.StatusOK, &d)
	// return captcha.WriteImage(c.Response().Writer(), d.CaptchaKey, 128, 44)

}

// 显示图片
func APICaptchaImage(c echo.Context) error {
	captchaKey := c.Param("key")
	// fmt.Println("CaptchaKey==>", captchaKey)
	// fmt.Println(c.Param("key"))
	return captcha.WriteImage(c.Response().Writer(), captchaKey, 240, 120)
	// img := captcha.NewImage(captchaKey, captcha.RandomDigits(4), 128, 44)
	// img.WriteTo(c.Response().Writer())
	// return nil

}

// 验证
func APICaptchaVerify(c echo.Context) error {
	key := c.FormValue("key")
	digits := c.FormValue("digits")
	rs := captcha.VerifyString(key, digits)
	return c.JSON(http.StatusOK, map[string]bool{"verify": rs})
}

//刷新
func APICaptchaReload(c echo.Context) error {
	key := c.Param("key")
	fmt.Println(key)
	rs := captcha.Reload(key)
	return c.JSON(http.StatusOK, map[string]bool{"isReload": rs})
}
