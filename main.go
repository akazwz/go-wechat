package main

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"log"
)

func main() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     "wxff07eecaebf3238f",
		AppSecret: "95591fc672b2f6973194865eb7cf8eeb",
		Cache:     memory,
	}

	mini := wc.GetMiniProgram(cfg)
	qr := mini.GetQRCode()
	code, err := qr.CreateWXAQRCode(qrcode.QRCoder{
		Page:      "",
		Path:      "./index",
		Width:     30,
		Scene:     "",
		AutoColor: false,
		LineColor: qrcode.Color{},
		IsHyaline: false,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(code))
}
