package lineBot

import "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"

type Service interface {
	GetBot() *messaging_api.MessagingApiAPI
	GetSecret() string
}
