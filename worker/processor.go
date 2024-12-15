package worker

import (
	"context"
	db "github.com/fsobh/auth/db/sqlc"
	"github.com/fsobh/mail"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

//This will pick up the tasks from the redis queue and process them

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
	mailer mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender) TaskProcessor {

	server := asynq.NewServer(redisOpt, asynq.Config{
		// These are the queues that we tell the processor to pick tasks up from with priority
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
		// Handling Queue Errors here :
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {

			log.Error().Err(err).Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("Error processing task")
		}),
		Logger: NewLogger(),
	},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	//use this to register each task with its handler function
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)

}
