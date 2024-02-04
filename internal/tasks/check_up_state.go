package tasks

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

type PayloadCheckUpState struct {
	url           string
	checkInterval int
	timeout       int
}

const TaskTypeCheckUpState = "task:check_up_state"

func (d *RedisTaskDistributor) DistributeTaskCheckUpState(ctx context.Context, payload PayloadCheckUpState, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TaskTypeCheckUpState, jsonPayload, opts...)
	info, err := d.client.Enqueue(task)
	if err != nil {
		return err
	}
	return nil
}
