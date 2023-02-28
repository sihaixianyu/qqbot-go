package service

import (
	"context"

	"github.com/solywsh/chatgpt"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/openapi"
)

var apiKey string
var chat *chatgpt.ChatGPT

func init() {
	logger = log.DefaultLogger

	apiKey = "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"
	chat = chatgpt.New(apiKey, "sihaixianyu", 0)
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
	chat.ChatContext.ResetConversation()

	answer, err := chat.Chat(content)
	if err != nil {
		logger.Error("GPT Chat failed: ", "err", err)
	}

	msg := buildSimpleMessage(data, answer)

	_, err = p.api.PostMessage(ctx, data.ChannelID, msg)
	if err != nil {
		logger.Error("Post message failed: ", "err", err)
	}

	chat.ChatContext.ResetConversation()
}

func (p *ATService) Chat(ctx context.Context, data *dto.WSATMessageData, content string) {
	var msg *dto.MessageToCreate
	var err error

	if content == "stop" {
		msg = buildSimpleMessage(data, "聊天已经终止！")
		chat.ChatContext.ResetConversation()
	} else {
		answer, err := chat.ChatWithContext(content)
		if err != nil {
			logger.Error("GPT Chat failed: ", "err", err)
		}

		msg = buildSimpleMessage(data, answer)
	}

	_, err = p.api.PostMessage(ctx, data.ChannelID, msg)
	if err != nil {
		logger.Error("Post message failed: ", "err", err)
	}
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
