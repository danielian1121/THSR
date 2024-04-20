package receiver

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strings"
)

var GinRouterSet = wire.NewSet(ProvideRoute, wire.Bind(new(Router), new(*GinRouter)))

type Receiver interface {
	GetRouteInfos() []ReceiverInfo
}

type ReceiverInfo struct {
	Method      string
	Path        string
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

func (r *ReceiverInfo) GetFlow() []gin.HandlerFunc {
	var flow []gin.HandlerFunc
	// Append specific middleware here
	flow = append(flow, r.Middlewares...)

	// Append handler
	flow = append(flow, r.Handler)
	return flow
}

type Router interface {
	RegisterAPI(engine *gin.Engine)
}

type GinRouter struct {
	routesInfo []ReceiverInfo
}

func (g *GinRouter) RegisterAPI(engine *gin.Engine) {
	for _, routeInfo := range g.routesInfo {
		engine.Handle(routeInfo.Method, g.PrefixPath(), routeInfo.GetFlow()...)
	}
}

func (r *GinRouter) PrefixPath() string {
	prefixes := []string{
		"v1",
	}
	return strings.Join(prefixes, "/")
}

func (g *GinRouter) generatePath(routeInfo ReceiverInfo) string {
	paths := []string{g.PrefixPath(), routeInfo.Path}
	return strings.Join(paths, "/")
}

func ProvideRoute(router ...Receiver) *GinRouter {
	return &GinRouter{
		routesInfo: extractRouteInfo(router...),
	}
}

func extractRouteInfo(receivers ...Receiver) []ReceiverInfo {
	var routeInfos []ReceiverInfo

	for _, receiver := range receivers {
		routeInfos = append(routeInfos, receiver.GetRouteInfos()...)
	}

	return routeInfos
}
