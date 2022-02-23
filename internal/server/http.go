package server

import (
	v1 "fiber-demo/api/helloworld/v1"
	"fiber-demo/internal/conf"
	"fiber-demo/internal/service"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/xhttp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *xhttp.Server {
	var opts = []xhttp.ServerOption{
		xhttp.Middleware(recover.New()),
		xhttp.FiberConfig(fiber.Config{
			JSONDecoder: encoding.GetCodec("json").Unmarshal,
			JSONEncoder: encoding.GetCodec("json").Marshal,
		}),
	}
	if c.Http.Network != "" {
		opts = append(opts, xhttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, xhttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, xhttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := xhttp.NewServer(opts...)
	v1.RegisterGreeterXHTTPServer(srv, greeter)
	return srv
}
