package main

import (
	"flag"
	"fmt"

	"backend-learning/blog/app/article/cmd/rpc/article"
	"backend-learning/blog/app/article/cmd/rpc/internal/config"
	"backend-learning/blog/app/article/cmd/rpc/internal/server"
	"backend-learning/blog/app/article/cmd/rpc/internal/svc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"manlu.org/tao/core/conf"
	"manlu.org/tao/core/service"
	"manlu.org/tao/zrpc"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewArticleServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		article.RegisterArticleServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
