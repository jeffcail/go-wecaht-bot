package handler

import (
	"log"
	"os"

	"github.com/eatmoreapple/openwechat"
)

func Handler(msg *openwechat.Message) {
	log.Printf("handler Received msg: %v", msg.Content)
	if msg.IsSendByGroup() {
		NewGroupMessageHandler().handle(msg)
		return
	}

	if msg.IsFriendAdd() {
		if os.Getenv("AutoPassFriendRequests") == "true" {
			_, err := msg.Agree("嗨！我是基于chatGPT的微信机器人,现在你可以向我提问任何问题哦...")
			if err != nil {
				log.Fatalf("add friend agress error: %v", err)
				return
			}
		}
	}
	NewUserMessageHandler().handle(msg)
}
