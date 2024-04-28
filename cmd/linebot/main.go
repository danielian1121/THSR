package main

import (
	"os"
	"path/filepath"

	"thsr/m/configs"
	"thsr/m/internal/injector"
	"thsr/m/server"
	"thsr/m/service/keyManager"

	"github.com/joho/godotenv"
)

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(path)
	godotenv.Load(filepath.Join(dir, ".env"))
}

func main() {
	manager := keyManager.New()
	configs.InitConfigs(manager)

	receiver := injector.BuildInjector()
	server.Init(receiver.Receiver)
}
