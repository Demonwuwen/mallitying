package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

//创建store
var store = base64Captcha.DefaultMemStore

//获取验证码
func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err

}

//验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id, VerifyValue)
	if store.Verify(id, VerifyValue, true) {
		fmt.Println("验证吗成功")
		return true
	} else {
		fmt.Println("验证吗失败")
		return false
	}
}
