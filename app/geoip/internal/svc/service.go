package svc

import (
	"simplex/pkg/jwt"
	"simplex/pkg/logx"
	"simplex/pkg/sid"
	"simplex/repository"
)

type Service struct {
	logger *logx.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *logx.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
