package main

import (
	"flag"

	"thsr/m/configs"
	"thsr/m/internal/injector"
	"thsr/m/server"
)

func main() {
	configs.InitConfigs()
	flag.Parse()

	receiver := injector.BuildInjector()
	server.Init(receiver.Receiver)
}
