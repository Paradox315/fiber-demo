// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fiber-demo/internal/biz"
	"fiber-demo/internal/conf"
	"fiber-demo/internal/data"
	"fiber-demo/internal/server"
	"fiber-demo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	xhttpServer := server.NewHTTPServer(confServer, greeterService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, xhttpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
