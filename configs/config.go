package configs

import (
	"context"
	"log"

	"thsr/m/service/keyManager"
)

var (
	C = new(config)
)

type config struct {
	LineBot LineBotConfig
}

type request struct {
	keys     []string
	callback func(map[string]string)
}

func InitConfigs() {
	manager := keyManager.New()
	register(manager)
}

func register(manager keyManager.Service) {
	var reqs []request
	reqs = append(reqs, lineBotFlags())

	getKeys(manager, reqs)
}

func getKeys(manager keyManager.Service, reqs []request) {
	keys := make([]string, 0, len(reqs))
	for _, req := range reqs {
		keys = append(keys, req.keys...)
	}

	m, err := manager.GetKeys(context.Background(), keys...)
	if err != nil {
		log.Fatalln("can't get secret value:", err)
		return
	}

	for _, req := range reqs {
		req.callback(m)
	}
}
