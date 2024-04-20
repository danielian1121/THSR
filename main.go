package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

var (
	defaultBot *messaging_api.MessagingApiAPI
	token      = os.Getenv("API_TOKEN")
	secret     = os.Getenv("SECRET_KEY")
)

func main() {
	bot, err := messaging_api.NewMessagingApiAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	defaultBot = bot
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Any("/callback", callback)

	// listen and serve on 0.0.0.0:8080
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func callback(c *gin.Context) {
	cb, err := webhook.ParseRequest(secret, c.Request)
	if err != nil {
		if errors.Is(err, webhook.ErrInvalidSignature) {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				if _, err = defaultBot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken, Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: message.Text,
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

				if _, err = defaultBot.ReplyMessage(
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
