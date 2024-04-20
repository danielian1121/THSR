package main

import (
	"thsr/m/configs"
	"thsr/m/internal/injector"
	"thsr/m/server"
)

func main() {
	configs.InitConfigs()

	receiver := injector.BuildInjector()
	server.Init(receiver.Receiver)
}
