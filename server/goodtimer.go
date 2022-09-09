package server

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
	"github.com/bhavpreet/goodTimer/parser"
	"github.com/bhavpreet/goodTimer/server/internal/config"
	"github.com/bhavpreet/goodTimer/server/internal/handler"
	"github.com/bhavpreet/goodTimer/server/internal/logic"
	"github.com/bhavpreet/goodTimer/server/internal/svc"
	"github.com/bhavpreet/goodTimer/server/internal/types"
	"github.com/bhavpreet/goodTimer/timer"
	"github.com/golang-collections/collections/stack"
	"github.com/timshannon/bolthold"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "server/etc/goodtimer-api.yaml", "the config file")
var store *bolthold.Store

func Serve() {
	flag.Parse()
	var err error

	var c config.Config
	conf.MustLoad(*configFile, &c)

	port := os.Getenv("PORT")
	if port != "" {
		var err error
		_port := c.Port
		c.Port, err = strconv.Atoi(port)
		if err != nil {
			c.Port = _port
		}
	}

	host := os.Getenv("HOST")
	if host != "" {
		c.Host = host
	}

	fmt.Printf("CFG = %+v", c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	ctx.Store, err = bolthold.Open("timer.db", 0666, nil)
	if err != nil {
		log.Printf("Unable to open the db..")
		return
	}
	defer ctx.Store.Close()
	store = ctx.Store

	// Create current
	_, err = logic.GetCurrent(ctx.Store)
	if err == bolthold.ErrNotFound {
		current := new(types.Current)
		current.ID = "current"
		err = ctx.Store.Insert("current", current)
		if err != nil {
			log.Printf("Unable to save current...")
			return
		}
	}

	// ctx.DB.Initialize(context.Background())
	handler.RegisterHandlers(server, ctx)

	// Timer
	// t, err := timer.NewTimer(&timer.TimerConfig{TimerType: c.TimerType}) // TODO
	t, err := timer.NewTimer(
		&timer.TimerConfig{TimerType: c.TimerType},
		processImpulse,
	)
	if err != nil {
		log.Fatalf("Error while initializing timer, err: %v", err)
	}
	go func() {
		for {
			err := t.Run()
			if err != nil {
				log.Printf("Error occurred while running the timer, err: %v", err)
				log.Println("Sleeping for 5 sec...")
				time.Sleep(5 * time.Second)
			}
		}
	}()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

var startStack = stack.New()

func processImpulse(ii *parser.Impulse) error {
	log.Println("Processing Impulse : ", ii)
	// check channel type / start or end
	if channelType, ok := timy2.ChannelType[ii.Channel]; ok {
		switch channelType {
		case timy2.START_IMPULSE:
			// Get current Round and Current START Bib
			startStack.Push(ii)
			current := new(types.Current)
			store.Get("current", current)
			if current.CurrentStartBib != "" {
				bib := new(types.Bib)
				err := store.Get(current.CurrentStartBib, bib)
				if err == nil {
					bib.StartTime = ii.Timestamp.Format(parser.DurationFormat)
					store.Update(current.CurrentStartBib, bib)

					// remove the current start
					current.CurrentStartBib = ""
					store.Update("current", current)
				}
			}
		case timy2.END_IMPULSE:
			// Get current Round and Current END Bib
			// If the BIB does not have a START, don't enter the END
			current := new(types.Current)
			store.Get("current", current)
			if current.CurrentEndBib != "" {
				bib := new(types.Bib)
				err := store.Get(current.CurrentEndBib, bib)
				if err == nil && bib.StartTime != "DNS" && bib.StartTime != "" {
					bib.EndTime = ii.Timestamp.Format(parser.DurationFormat)
					store.Update(current.CurrentEndBib, bib)

					// remove the current start
					current.CurrentEndBib = ""
					store.Update("current", current)
				}
			}

			if start := startStack.Peek(); start == nil {
				println("False Start", ii.Channel)
			} else {
				_start, _ := startStack.Pop().(*parser.Impulse)
				var t parser.Timespan
				t = parser.Timespan(ii.Timestamp.Sub(_start.Timestamp))
				log.Println("FINISH:", t.Format(parser.DurationFormat))
			}
		}
	} else {
		println("Unknown channel type " + ii.Channel)
	}
	return nil
}
