package svc

import (
	"github.com/bhavpreet/goodTimer/server/internal/config"
	bh "github.com/timshannon/bolthold"
)

type ServiceContext struct {
	*bh.Store
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
