package redis

import (
	"fmt"
	"github.com/spiral/roadrunner/service"
	"github.com/spiral/roadrunner/service/rpc"
)

// ID defines public service name.
const ID = "redis"

type Config struct {
	Addr string
	Password string
	DB int
}

func (c *Config) Hydrate(cfg service.Config) error {
	return cfg.Unmarshal(&c)
}

// Service manages even broadcasting and websocket interface.
type Service struct {}

func (s *Service) Init(r *rpc.Service, cfg *Config) (ok bool, err error) {
	fmt.Println(cfg)
	return true, nil
}