package main

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"log"
)

type LoginData struct {
	Code    string `json:"code"`
	Encrypt string `json:"encrypt"`
	Iv      string `json:"iv"`
}

func main() {
	SendMsg()
	/*r := gin.Default()
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
	r.Run()*/
}

func SendMsg() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     "wx888a82e958faaaaa",
		AppSecret: "67b8b7ae4d304f3bc44145f820c1ba24",
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	sub := mini.GetSubscribe()
	data := make(map[string]*subscribe.DataItem)
	data["phrase1"] = &subscribe.DataItem{
		Value: "test",
		Color: "",
	}
	data["phrase1"] = &subscribe.DataItem{
		Value: "吃饭了干嘛吗",
		Color: "",
	}
	data["date2"] = &subscribe.DataItem{
		Value: "2019-12-11 11:00:00",
		Color: "",
	}
	data["phrase3"] = &subscribe.DataItem{
		Value: "点击查看",
		Color: "",
	}
	msg := &subscribe.Message{
		ToUser:     "ov3y85KApNeZY_-yCtmy-5X--Un8",
		TemplateID: "XV16ZyG6Af_gG8D4qg7M17Fw23m_zYWNo689XpJKYQE",
		Data:       data,
	}
	err := sub.Send(msg)
	if err != nil {
		log.Println("send error")
		log.Println(err)
		return
	}
}
