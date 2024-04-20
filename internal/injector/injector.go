package injector

import (
	"github.com/google/wire"

	"thsr/m/server/receiver"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Receiver receiver.Router
}
