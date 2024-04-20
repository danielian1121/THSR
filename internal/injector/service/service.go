package service

import (
	"github.com/google/wire"

	"thsr/m/service/lineBot"
)

var ServiceSet = wire.NewSet(lineBot.New)
