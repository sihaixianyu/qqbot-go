package main

import (
	"strings"

	"github.com/sihaixianyu/qqbot-go/service"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
)

// ReadyHandler 自定义 ReadyHandler 感知连接成功事件
func ReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		logger.Info("Ready event receive: ", "WSReadyData", data)
	}
}

// ErrorNotifyHandler 自定义 ErrorNotifyHandler 感知错误 
func ErrorNotifyHandler() event.ErrorNotifyHandler {
	return func(err error) {
		logger.Error("Error notify receive: ", "err", err)
	}
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler(api openapi.OpenAPI) event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {

		input := strings.ToLower(message.ETLInput(data.Content))
		atService := service.NewATService(api)

		return atService.Process(input, data)
	}
}
