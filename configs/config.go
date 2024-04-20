package configs

var (
	C = new(config)
)

type config struct {
	LineBot LineBotConfig
}

func InitConfigs() {
	registerFlags()
}

func registerFlags() {
	lineBotFlags()
}
