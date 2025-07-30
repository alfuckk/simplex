package task

import (
	"simplex/pkg/jwt"
	"simplex/pkg/logx"
	"simplex/pkg/sid"
	"simplex/repository"
)

type Task struct {
	logger *logx.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *logx.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
