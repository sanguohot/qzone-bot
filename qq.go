package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)
func main() {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName":                      "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("G:/Program Files (ssd)/browsedriver/chromedriver", 8080, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()
	// 调起chrome浏览器
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 8080))
	if err != nil {
		panic(err)
	}
	if err = wd.Get("https://i.qq.com/"); err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	err = wd.SwitchFrame("login_frame")
	if err != nil {
		panic(err)
	}
	switcher, err := wd.FindElement(selenium.ByID, "switcher_plogin")
	if err != nil {
		panic(err)
	}
	err = switcher.Click()
	if err != nil {
		panic(err)
	}
	username, err := wd.FindElement(selenium.ByID, "u")
	if err != nil {
		panic(err)
	}
	password, err := wd.FindElement(selenium.ByID, "p")
	if err != nil {
		panic(err)
	}
	err = username.SendKeys("281406674")
	if err != nil {
		panic(err)
	}
	err = password.SendKeys("hw535431")
	if err != nil {
		panic(err)
	}
	button, err := wd.FindElement(selenium.ByID, "login_button")
	if err != nil {
		panic(err)
	}
	err = button.Click()
	if err != nil {
		panic(err)
	}
	//go func() {
	//	for{
	//		if msg, err := wd.FindElement("id", "err_m"); err != nil {
	//			log.Println(err.Error())
	//			panic(err)
	//		}else {
	//			log.Println(msg.Text())
	//		}
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//}()
	time.Sleep(time.Second * 1)
	menuContainer, err := wd.FindElement(selenium.ByID, "menuContainer")
	if err != nil {
		panic(err)
	}
	photo, err := menuContainer.FindElement(selenium.ByLinkText, "相册")
	if err != nil {
		panic(err)
	}
	fmt.Println(photo.Text())
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
}
