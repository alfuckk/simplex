package task

import (
	"context"
	"simplex/app/geoip/internal/repo"
)

type UserTask interface {
	CheckUser(ctx context.Context) error
}

func NewUserTask(
	task *Task,
	userRepo repo.UserRepository,
) UserTask {
	return &userTask{
		userRepo: userRepo,
		Task:     task,
	}
}

type userTask struct {
	userRepo repo.UserRepository
	*Task
}

func (t userTask) CheckUser(ctx context.Context) error {
	// do something
	t.logger.Info("CheckUser")
	return nil
}
