package redis

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spiral/roadrunner/service"
	"github.com/spiral/roadrunner/service/rpc"
	"testing"
	"time"
)

var testString = `{"some":"info","object":{"name":"Vasya"}}`
var rpcPort = 6001
var rdSvc *Service

type testCfg struct {
	rpc    string
	redis  string
	target string
}

func (cfg *testCfg) Get(name string) service.Config {
	if name == ID {
		return &testCfg{target: cfg.redis}
	}

	if name == rpc.ID {
		return &testCfg{target: cfg.rpc}
	}

	return nil
}

func (cfg *testCfg) Unmarshal(out interface{}) error {
	return json.Unmarshal([]byte(cfg.target), out)
}

func init() {
	rdSvc, _, _ = setup(`{"redis":{"addr":"localhost:6379"}}`)
}

func setup(cfg string) (*Service, *rpc.Service, service.Container) {
	logger, _ := test.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	c := service.NewContainer(logger)
	c.Register(rpc.ID, &rpc.Service{})
	c.Register(ID, &Service{})

	err := c.Init(&testCfg{
		redis: cfg,
		rpc:   fmt.Sprintf(`{"enable": true, "listen":"tcp://:%v"}`, rpcPort),
	})

	rpcPort++

	if err != nil {
		panic(err)
	}

	go func() {
		err := c.Serve()
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)

	svc, _ := c.Get(ID)
	redisService := svc.(*Service)

	rpcSvc, _ := c.Get(rpc.ID)
	rpcService := rpcSvc.(*rpc.Service)

	return redisService, rpcService, c
}

func TestService_Init(t *testing.T) {

}
