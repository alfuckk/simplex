package job

import (
	"context"
	"simplex/app/auth/internal/repo"
	"time"
)

type UserJob interface {
	KafkaConsumer(ctx context.Context) error
}

func NewUserJob(
	job *Job,
	userRepo repo.UserRepository,
) UserJob {
	return &userJob{
		userRepo: userRepo,
		Job:      job,
	}
}

type userJob struct {
	userRepo repo.UserRepository
	*Job
}

func (t userJob) KafkaConsumer(ctx context.Context) error {
	// do something
	for {
		t.logger.Info("KafkaConsumer")
		time.Sleep(time.Second * 5)
	}
}
