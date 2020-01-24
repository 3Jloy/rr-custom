package redis

import (
	"encoding/json"
	"time"
)

type Message struct {
	Key        string          `json:"key"`
	Value      json.RawMessage `json:"value"`
	Expiration string          `json:"expiration"`
}

type rpcService struct {
	svc *Service
}

func (s *rpcService) Set(input Message, output *bool) error {
	expiration, err := time.ParseDuration(input.Expiration)
	if err != nil {
		*output = false
		return err
	}

	var value []byte
	value, _ = input.Value.MarshalJSON()

	_, err = s.svc.redisClient.Set(input.Key, value, expiration).Result()
	if err != nil {
		*output = false
		return err
	}

	*output = true
	return nil
}

func (s *rpcService) Get(input Message, output *string) error {
	result, err := s.svc.redisClient.Get(input.Key).Result()
	*output = result
	return err
}

func (s *rpcService) Del(input []string, output *bool) error {
	if _, err := s.svc.redisClient.Del(input...).Result(); err != nil {
		return err
	}

	*output = true
	return nil
}
