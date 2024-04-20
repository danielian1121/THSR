package api

import (
	"github.com/google/wire"

	"thsr/m/server/receiver"
)

var RouterSet = wire.NewSet(receiver.GinRouterSet)
