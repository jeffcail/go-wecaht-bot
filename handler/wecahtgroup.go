package handler

import (
	"log"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/jeffcail/go-wechat-bot/common/chatGpt"
)

type GroupMessageHandler struct{}

func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.replyGroupText(msg)
	}
	return nil
}

func NewGroupMessageHandler() *GroupMessageHandler {
	return &GroupMessageHandler{}
}

func (g *GroupMessageHandler) replyGroupText(msg *openwechat.Message) error {
	// 获取消息的发送者
	sender, err := msg.Sender()
	if err != nil {
		return err
	}
	group := openwechat.Group{sender}
	log.Printf("received group %v msg : %v", group.NickName, msg.Content)

	// 判断消息是否为@消息
	if !msg.IsAt() {
		return nil
	}

	replaceText := "@" + sender.Self.NickName
	reqestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	reply, err := chatGpt.Completions(reqestText)
	if err != nil {
		log.Printf("gtp request error: %v\n", err)
		msg.ReplyText("哇！🐂👃")
		return err
	}
	if reply == "" {
		return nil
	}

	// 获取消息在群里面的发送者
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		log.Printf("get sender in group error :%v \n", err)
		return err
	}

	reply = strings.Trim(strings.TrimSpace(reply), "\n")
	atText := "@" + groupSender.NickName
	replyText := atText + reply
	_, err = msg.ReplyText(replyText)
	if err != nil {
		log.Printf("reply group error: %v \n", err)
		return err
	}
	return nil
}
