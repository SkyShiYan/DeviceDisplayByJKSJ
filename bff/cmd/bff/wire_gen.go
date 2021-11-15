// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"bff/internal/biz"
	"bff/internal/conf"
	"bff/internal/data"
	"bff/internal/server"
	"bff/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	deviceRepo := data.NewBffRepo(dataData, logger)
	bffUsecase := biz.NewBffUsecase(deviceRepo, logger)
	bffService := service.NewBffService(bffUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, bffService, logger)
	grpcServer := server.NewGRPCServer(confServer, bffService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}