package handler

import (
	"log"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/jeffcail/go-wechat-bot/common/chatGpt"
)

type UserMessageHandler struct{}

func (u *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return u.replyUserText(msg)
	}
	return nil
}

func NewUserMessageHandler() *UserMessageHandler {
	return &UserMessageHandler{}
}

func (u *UserMessageHandler) replyUserText(msg *openwechat.Message) error {
	// 获取消息的发送者
	sender, err := msg.Sender()
	if err != nil {
		return err
	}
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)

	requestText := strings.Trim(strings.TrimSpace(msg.Content), "\n")
	reply, er := chatGpt.Completions(requestText)
	if er != nil {
		log.Printf("gtp request error: %v\n", er)
		msg.ReplyText("哇！🐂👃")
		return er
	}
	if reply == "" {
		return nil
	}
	reply = strings.Trim(strings.TrimSpace(reply), "\n")
	_, err = msg.ReplyText(reply)
	if err != nil {
		log.Printf("reply user error: %v \n", err)
		return err
	}
	return nil
}
