package redis

import (
	"fmt"
	"github.com/spiral/roadrunner/service/rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rpcSvc *rpc.Service

func init() {
	_, rpcSvc, _ = setup(`{"redis":{"addr":"localhost:6379"}}`)
}

func TestRpcService_Set(t *testing.T) {
	rpcClient, err := rpcSvc.Client()
	assert.NoError(t, err)

	rpcMethod := "redis.Set"

	output := false
	assert.NoError(t, rpcClient.Call(
		rpcMethod,
		Message{
			Key:        "1",
			Value:      testString,
			Expiration: 30,
		},
		&output,
	))
	assert.True(t, output)

	rpcClient, err = rpcSvc.Client()
	assert.NoError(t, err)

	output = false
	err = rpcClient.Call(
		rpcMethod,
		Message{
			Key:        "2",
			Value:      testString,
			Expiration: 30,
		},
		&output,
	)

	assert.False(t, output)
	assert.Equal(t, fmt.Errorf("reading body time: missing unit in duration 5"), err)

	rpcClient, err = rpcSvc.Client()
	assert.NoError(t, err)
}

func TestRpcService_Get(t *testing.T) {
	rpcClient, err := rpcSvc.Client()
	assert.NoError(t, err)

	output := ""
	assert.NoError(t, rpcClient.Call(
		"redis.Get",
		Message{
			Key: "1",
		},
		&output,
	))

	assert.Equal(t, testString, output)

	rpcClient, err = rpcSvc.Client()
	assert.NoError(t, err)

	output = ""
	err = rpcClient.Call(
		"redis.Get",
		Message{
			Key: "99999999999",
		},
		&output,
	)

	assert.Equal(t, "", output)
	assert.Equal(t, fmt.Errorf("reading body redis: nil"), err)
}

func TestRpcService_Del(t *testing.T) {
	rpcClient, err := rpcSvc.Client()
	assert.NoError(t, err)

	output := false
	assert.NoError(t, rpcClient.Call(
		"redis.Del",
		[]string{"1"},
		&output,
	))

	assert.True(t, output)

	getOutput := ""
	redisGetErr := rpcClient.Call("redis.Get",
		Message{
			Key: "1",
		},
		&getOutput,
	)
	assert.Error(t, redisGetErr)

}
