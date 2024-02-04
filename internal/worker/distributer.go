package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskCheckUpState(ctx context.Context, payload PayloadCheckUpState, opts ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func newRedisTaskDistributer(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
