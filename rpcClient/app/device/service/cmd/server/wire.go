// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"spaco-1103/app/device/service/internal/biz"
	"spaco-1103/app/device/service/internal/conf"
	"spaco-1103/app/device/service/internal/data"
	"spaco-1103/app/device/service/internal/server"
	"spaco-1103/app/device/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
