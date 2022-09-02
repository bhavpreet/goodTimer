package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bhavpreet/goodTimer/server/db"
	"github.com/bhavpreet/goodTimer/server/internal/config"
	"github.com/bhavpreet/goodTimer/server/internal/handler"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/timer"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/goodtimer-api.yaml", "the config file")

func Serve() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	port := os.Getenv("PORT")
	if port != "" {
		var err error
		c.Port, err = strconv.Atoi(port)
		if err != nil {
			// do nothing
		}
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	ctx.DB = db.NewDefaultDB()
	ctx.DB.Initialize(context.Background())
	handler.RegisterHandlers(server, ctx)

	// Timer
	log.Printf("cfg = %+v", c)
	// t, err := timer.NewTimer(&timer.TimerConfig{TimerType: c.TimerType}) // TODO
	t, err := timer.NewTimer(&timer.TimerConfig{TimerType: timer.SIMULATOR})
	if err != nil {
		log.Fatalf("Error while initializing timer, err: %v", err)
	}
	go func() {
		for {
			err := t.Run()
			if err != nil {
				log.Printf("Error occurred  while running the timer, err: %v", err)
				log.Println("Sleeping for 5 sec...")
				time.Sleep(5 * time.Second)
			}
		}
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
