package job

import (
	"simplex/pkg/jwt"
	"simplex/pkg/logx"
	"simplex/pkg/sid"
	"simplex/repository"
)

type Job struct {
	logger *logx.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *logx.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
