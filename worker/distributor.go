package worker

import (
	"context"
	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		option ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

// NewRedisTaskDistributor we return a new redis task distributor interface to force the RedisTaskDistributor to implement it
func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	//create a new redis client

	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{client: client}
}
