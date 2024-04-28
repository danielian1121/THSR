package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"thsr/m/configs"
	"thsr/m/internal/injector"
	"thsr/m/server"
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
	configs.InitConfigs()

	receiver := injector.BuildInjector()
	server.Init(receiver.Receiver)
}
