package worker

import (
	"context"
	"encoding/json"
	"fmt"
	db "github.com/fsobh/auth/db/sqlc"
	"github.com/fsobh/auth/util"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "TaskSendVerifyEmail"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option) error {

	//serialize payload struct into bytes

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf(
			"failed to marshal payload: %w",
			err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)

	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf(
			"failed to enqueue task: %w",
			err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("task enqueued")

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {

	var payload PayloadSendVerifyEmail
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf(
			"failed to unmarshal payload: %w",
			asynq.SkipRetry)
	}
	user, err := processor.store.GetUser(ctx, payload.Username)

	if err != nil {
		// we don't need this - even if the user wasn't found, we will let it retry (incase the db is slow for whatever reason) - we don't want a user to get created without an mail being sent if delays occur for the tx commit
		//if errors.Is(err, sql.ErrNoRows) {
		//	return fmt.Errorf("user not found %w", asynq.SkipRetry)
		//}
		return fmt.Errorf("failed to get user: %w", err)
	}

	userVerifyEmail := db.CreatVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	}
	verifyEmail, err := processor.store.CreatVerifyEmail(ctx, userVerifyEmail) // make record of it in db

	if err != nil {
		return fmt.Errorf("failed to create verify mail: %w", err)
	}

	subject := "Verify your email"
	verifyEmailURL := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s", verifyEmail.ID, verifyEmail.SecretCode)
	contentHtml := fmt.Sprintf("<h1>Verify your email</h1><p>Please click <a href='%s'>here</a> to verify your email</p>", verifyEmailURL)
	to := []string{user.Email}

	// Send Email here
	err = processor.mailer.SendMail(subject, contentHtml, to, nil, nil, nil)

	if err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("mail", user.Email).Msg("processed task")

	return nil
}
