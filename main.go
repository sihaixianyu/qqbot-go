package main

import (
	"context"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

type Config struct {
	AppID uint64 `toml:"app_id"`
	Token string `toml:"token"`
}

var logger log.Logger

func init() {
	logger = log.DefaultLogger
}

func main() {
	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Error(err)
	}

	botToken := token.BotToken(config.AppID, config.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)

	ctx := context.Background()
	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Error(err)
	}

	intent := websocket.RegisterHandlers(
		// 连接成功回调
		ReadyHandler(),
		// 连接错误回调
		ErrorNotifyHandler(),
		// @机器人事件
		ATMessageEventHandler(api),
	)

	if err = botgo.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		log.Error(err)
	}
}
