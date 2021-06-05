package server

import (
	v1 "blog-demo/api/blog/v1"
	"blog-demo/internal/conf"
	"blog-demo/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, tracer trace.TracerProvider ,blog *service.BlogService) *http.Server {
	var opts []http.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		recovery.Recovery(),
		tracing.Server(tracing.WithTracerProvider(tracer)),
		logging.Server(log.DefaultLogger),
		validate.Validator(),
	)
	srv.HandlePrefix("/", v1.NewBlogHandler(blog, m))
	return srv
}
