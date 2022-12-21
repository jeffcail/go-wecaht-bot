# go-wechat-bot
玩一玩cahtGPT对接微信.实现 自动通过好友申请、用户私聊回复、群聊艾特回复

### wechat SDK
```shell
go get github.com/eatmoreapple/openwechat
```

### 配置
```shell
export OpenApiKey="" # chatGPT api key
export AutoPassFriendRequests = "true"
```

### 使用
```shell
git clone https://github.com/jeffcail/go-wechat-bot.git


åcd go-wechat-bot

go run main.go
```