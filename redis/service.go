package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spiral/roadrunner/service"
	"github.com/spiral/roadrunner/service/rpc"
)

// ID defines public service name.
const ID = "redis"

type Config struct {
	Addr     string
	Password string
	DB       int
}

func (c *Config) Hydrate(cfg service.Config) error {
	return cfg.Unmarshal(&c)
}

// Service manages even broadcasting and websocket interface.
type Service struct {
	redisClient redis.UniversalClient
}

func (s *Service) Init(r *rpc.Service, cfg *Config) (ok bool, err error) {
	fmt.Println(cfg)
	s.redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := r.Register(ID, &rpcService{svc: s}); err != nil {
		return false, err
	}

	return true, nil
}

type Message struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type rpcService struct {
	svc *Service
}

func (s *rpcService) Set(input Message, ok *bool) error {
	fmt.Println(s)
	value, _ := input.Value.MarshalJSON()
	fmt.Println(input, string(value))

	return nil
}
