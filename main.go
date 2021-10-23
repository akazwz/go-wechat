package main

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"log"
)

type LoginData struct {
	Code    string `json:"code"`
	Encrypt string `json:"encrypt"`
	Iv      string `json:"iv"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(ctx *gin.Context) {
		var data LoginData
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			log.Println("获取数据失败")
		}
		wc := wechat.NewWechat()
		memory := cache.NewMemory()
		cfg := &config.Config{
			AppID:     "wx888a82e958faaaaa",
			AppSecret: "67b8b7ae4d304f3bc44145f820c1ba24",
			Cache:     memory,
		}
		mini := wc.GetMiniProgram(cfg)
		auth := mini.GetAuth()
		session, err := auth.Code2Session(data.Code)
		if err != nil {
			log.Println("获取 session 错误")
			return
		}
		log.Println("uni:" + session.UnionID)
		log.Println("open:" + session.OpenID)
		encryptor := mini.GetEncryptor()
		plainData, err := encryptor.Decrypt(session.SessionKey, data.Encrypt, data.Iv)
		if err != nil {
			log.Println("解密数据错误")
			return
		}
		log.Println(len(plainData.PhoneNumber))
	})
	r.Run()
}
