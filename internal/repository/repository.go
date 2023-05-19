package repository

import "sync"

type Repository interface {
	Job
}

type repository struct {
	storage sync.Map
}

func NewRepository() Repository {

	return &repository{storage: sync.Map{}}
}
