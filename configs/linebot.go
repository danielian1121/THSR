package configs

const (
	token  = "linebot/CHANNEL_ACCESS_TOKEN"
	secret = "linebot/CHANNEL_SECRET"
)

type LineBotConfig struct {
	Token  string
	Secret string
}

func lineBotFlags() request {
	return request{
		keys: []string{token, secret},
		callback: func(m map[string]string) {
			C.LineBot.Token = m[token]
			C.LineBot.Secret = m[secret]
		},
	}
}
