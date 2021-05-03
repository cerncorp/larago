package services

import (
	"github.com/cerncorp/larago/app/models"
	"github.com/cerncorp/larago/config"
)

type CallableService interface {
	// methods
	CallService(name string, age string) (string, error)
	Save(video models.Video) error
	FindAll() []models.Video
}

type videoService struct {
	// attributes
	videos []models.Video
}

func NewCallableService() CallableService {
	return &videoService{
		videos: []models.Video{},
	}
}

func (s *videoService) CallService(name string, age string) (string, error) {
	result := name + age
	config.Log.Info(result)
	config.Log.Error(result)
	config.Log.Warning(result)
	return result, nil
}

func (s *videoService) Save(video models.Video) error {
	s.videos = append(s.videos, video)
	return nil
}

func (s *videoService) FindAll() []models.Video {
	return s.videos
}
