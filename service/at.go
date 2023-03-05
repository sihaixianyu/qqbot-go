package service

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/openapi"
)

var apiKey string

func init() {
	logger = log.DefaultLogger
	apiKey = "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"
}

type ATService struct {
	api openapi.OpenAPI
}

func NewATService(api openapi.OpenAPI) ATService {
	return ATService{
		api: api,
	}
}

// Process is a function to process message
func (p ATService) Process(input string, data *dto.WSATMessageData) error {
	ctx := context.Background()
	cmd := message.ParseCommand(input)

	switch cmd.Cmd {
	case "/ask":
		p.Ask(ctx, data, cmd.Content)
	case "/chat":
		p.Chat(ctx, data, cmd.Content)
	default:
		content := "旅行者有什么事情吗" + message.Emoji(32)
		msg := buildSimpleMessage(data, content)

		_, err := p.api.PostMessage(ctx, data.ChannelID, msg)
		if err != nil {
			logger.Error("Default reply failed: ", "err", err)
		}
	}

	return nil
}

func (p *ATService) Ask(ctx context.Context, data *dto.WSATMessageData, content string) {

}

func (p *ATService) Chat(ctx context.Context, data *dto.WSATMessageData, content string) {

}

func buildSimpleMessage(data *dto.WSATMessageData, content string) *dto.MessageToCreate {
	msg := &dto.MessageToCreate{
		Content: content,
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}

	return msg
}
