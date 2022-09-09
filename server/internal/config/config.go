package config

import (
	"github.com/bhavpreet/goodTimer/timer"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	TimerType timer.TimerType
}
