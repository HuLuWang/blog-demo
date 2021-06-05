// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog-demo/internal/biz"
	"blog-demo/internal/conf"
	"blog-demo/internal/data"
	"blog-demo/internal/server"
	"blog-demo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
