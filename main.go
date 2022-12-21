package main

import (
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/jeffcail/go-wechat-bot/handler"
)

func main() {
	// bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	if msg.IsText() && msg.Content == "ping" {
	//		msg.ReplyText("pong")
	//	}
	//}
	bot.MessageHandler = handler.Handler

	// 注册二维码登陆
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	storage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	err := bot.HotLogin(storage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v\n", err)
			return
		}
	}
	bot.Block()
}
