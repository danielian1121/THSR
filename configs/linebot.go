package configs

import "flag"

type LineBotConfig struct {
	Token  string
	Secret string
}

func lineBotFlags() {
	flag.StringVar(&C.LineBot.Token, "line_bot_token", "", "line bot token")
	flag.StringVar(&C.LineBot.Secret, "line_bot_channel_secret", "", "channel secret")
}
