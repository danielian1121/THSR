package webhook

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"

	"thsr/m/server/receiver"
	"thsr/m/service/lineBot"
)

type impl struct {
	lineBot lineBot.Service
}

func ProvideReceiver(
	lineBot lineBot.Service,
) Receiver {
	h := impl{
		lineBot: lineBot,
	}

	return &h
}

func (im *impl) handleLineWebhook(c *gin.Context) {
	cb, err := webhook.ParseRequest(im.lineBot.GetSecret(), c.Request)
	if err != nil {
		if errors.Is(err, webhook.ErrInvalidSignature) {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	bot := im.lineBot.GetBot()
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							&messaging_api.TemplateMessage{
								AltText: message.Text,
								Template: &messaging_api.ConfirmTemplate{
									Text: "liff",
									Actions: []messaging_api.ActionInterface{
										&messaging_api.UriAction{
											Label: "Yes",
											Uri:   "https://liff.line.me/2004698296-Vv4bvpZ7",
										},
										&messaging_api.MessageAction{
											Label: "No",
											Text:  "No!",
										},
									},
								},
							},
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent text reply.")
				}
			case webhook.StickerMessageContent:
				replyMessage := fmt.Sprintf("sticker id is %s, stickerResourceType is %s", message.StickerId, message.StickerResourceType)

				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: replyMessage,
							},
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent sticker reply.")
				}
			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}

type liffVerifyParm struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

func (im *impl) liffVerify(c *gin.Context) {
	var param liffVerifyParm
	if err := c.BindJSON(&param); err != nil {
		c.JSON(420, gin.H{})
		return
	}
	log.Println(param.AccessToken)
	c.Redirect(http.StatusFound, fmt.Sprintf("https://api.line.me/oauth2/v2.1/verify?access_token=%s", param.AccessToken))
}

func (im *impl) GetRouteInfos() []receiver.ReceiverInfo {
	return []receiver.ReceiverInfo{
		{
			Method:  http.MethodPost,
			Path:    "/webhook",
			Handler: im.handleLineWebhook,
		},
		{
			Method:  http.MethodPost,
			Path:    "/liff/verify",
			Handler: im.liffVerify,
		},
	}
}
