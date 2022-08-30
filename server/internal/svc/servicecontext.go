package svc

import (
	"github.com/bhavpreet/goodTimer/server/db"
	"github.com/bhavpreet/goodTimer/server/internal/config"
)

type ServiceContext struct {
	db.DB
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
