package service

import (
	"simple-proxy/internal/repository"
	"simple-proxy/pkg/logger"
)

type Service interface {
	Job
}

type service struct {
	logger     logger.Logger
	repository repository.Repository
}

func NewService(logger logger.Logger, repository repository.Repository) Service {
	return &service{logger: logger, repository: repository}
}
