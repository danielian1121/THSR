//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"thsr/m/internal/injector/api"
	"thsr/m/internal/injector/service"
)

func BuildInjector() *Injector {
	wire.Build(
		api.RouterSet,
		api.ReceiverSet,
		api.ProvideReceiverList,

		service.ServiceSet,

		InjectorSet,
	)

	return new(Injector)
}
