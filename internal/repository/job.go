package repository

import "simple-proxy/internal/model"

type Job interface {
	CreateJob(job model.Job)
	GetJob(id string) (*model.Job, bool)
}

func (r *repository) CreateJob(job model.Job) {
	r.storage.Store(job.ID, &job)
}

func (r *repository) GetJob(id string) (*model.Job, bool) {
	val, ok := r.storage.Load(id)
	if !ok {
		return nil, false
	}

	return val.(*model.Job), true
}
