package service

import (
	"github.com/google/uuid"
	"github.com/kaolnwza/spider-cat-microsvc/spider-service/model"
	repository "github.com/kaolnwza/spider-cat-microsvc/spider-service/repositories"
)

type spiderSvc struct {
	spiderRepo repository.SpiderRepository
}

type SpiderService interface {
	GetSpidersService() ([]model.Spider, error)
	GetSpiderByUUIDService(uuid uuid.UUID) (model.Spider, error)
	AttackSpiderService(damage int) error
}

func NewSpiderService(spiderRepo repository.SpiderRepository) spiderSvc {
	return spiderSvc{spiderRepo: spiderRepo}
}

func (s spiderSvc) GetSpidersService() ([]model.Spider, error) {
	spider, err := s.spiderRepo.GetSpiders()
	if err != nil {
		return nil, err
	}

	return spider, nil
}

func (s spiderSvc) GetSpiderByUUIDService(uuid uuid.UUID) (model.Spider, error) {
	return model.Spider{}, nil
}

func (s spiderSvc) AttackSpiderService(damage int) error {
	return nil
}
