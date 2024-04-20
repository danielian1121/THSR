package lineBot

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"thsr/m/configs"
)

type impl struct {
	bot    *messaging_api.MessagingApiAPI
	secret string
}

func New() Service {
	bot, err := messaging_api.NewMessagingApiAPI(configs.C.LineBot.Token)
	if err != nil {
		panic(err)
	}

	s := &impl{
		bot:    bot,
		secret: configs.C.LineBot.Secret,
	}
	return s
}

func (im *impl) GetBot() *messaging_api.MessagingApiAPI {
	return im.bot
}

func (im *impl) GetSecret() string {
	return im.secret
}
