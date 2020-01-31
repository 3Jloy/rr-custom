package redis

import (
	"time"
)

type Message struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int    `json:"expiration"`
}

type rpcService struct {
	svc *Service
}

func (s *rpcService) Set(input Message, output *bool) error {
	expiration := time.Second * time.Duration(input.Expiration)
	_, err := s.svc.redisClient.Set(input.Key, input.Value, expiration).Result()
	*output = err != nil
	return err
}

func (s *rpcService) Get(input Message, output *string) error {
	result, err := s.svc.redisClient.Get(input.Key).Result()
	*output = result
	return err
}

func (s *rpcService) Del(input []string, output *bool) error {
	_, err := s.svc.redisClient.Del(input...).Result()
	*output = err != nil
	return err
}
