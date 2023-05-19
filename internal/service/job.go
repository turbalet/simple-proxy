package service

import (
	"github.com/google/uuid"
	"net/http"
	"simple-proxy/internal/model"
	"simple-proxy/internal/model/dto"
	"time"
)

type Job interface {
	GetJob(id string) (*model.Job, error)
	CreateJob(in dto.JobRequest) (*dto.JobResponse, error)
}

func (s *service) GetJob(id string) (*model.Job, error) {
	job, has := s.repository.GetJob(id)
	if !has {
		return nil, ErrJobNotFound
	}

	return job, nil
}

func (s *service) CreateJob(in dto.JobRequest) (*dto.JobResponse, error) {
	validHeader := make(map[string][]string)
	for k, v := range in.Headers {
		validHeader[k] = append(validHeader[k], v)
	}

	removeHopHeaders(validHeader)

	request, err := http.NewRequest(in.Method, in.URL, nil)
	if err != nil {
		s.logger.Errorf("err creating request: %v", err)
		return nil, err
	}
	request.Header = validHeader

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		s.logger.Errorf("err sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	removeHopHeaders(resp.Header)

	id := uuid.New().String()

	s.repository.CreateJob(model.Job{
		ID:            id,
		Status:        resp.StatusCode,
		URL:           in.URL,
		ContentLength: resp.ContentLength,
		ReqMethod:     in.Method,
		ReqHeaders:    validHeader,
		RespHeaders:   resp.Header,
	})

	return &dto.JobResponse{
		ID:      id,
		Headers: resp.Header,
		Length:  resp.ContentLength,
		Status:  resp.StatusCode,
	}, nil
}
