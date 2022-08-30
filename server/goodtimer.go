package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bhavpreet/goodTimer/server/db"
	"github.com/bhavpreet/goodTimer/server/internal/config"
	"github.com/bhavpreet/goodTimer/server/internal/handler"
	"github.com/bhavpreet/goodTimer/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/goodtimer-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	ctx.DB = db.NewDefaultDB()
	ctx.DB.Initialize(context.Background())
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
