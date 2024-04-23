package configs

import (
	"flag"
	"os"
)

type LineBotConfig struct {
	Token  string
	Secret string
}

func lineBotFlags() {
	flag.StringVar(&C.LineBot.Token, "line_bot_token", os.Getenv("LINE_BOT_TOKEN"), "line bot token")
	flag.StringVar(&C.LineBot.Secret, "line_bot_channel_secret", os.Getenv("LINE_BOT_CHANNEL_SECRET"), "channel secret")
}
