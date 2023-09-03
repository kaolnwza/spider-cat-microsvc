package repository

import (
	"github.com/google/uuid"
	"github.com/kaolnwza/spider-cat-microsvc/spider-service/model"
)

type spiderRepo struct{}

type SpiderRepository interface {
	GetSpiders() ([]model.Spider, error)
	GetSpiderByUUID(uuid uuid.UUID) (model.Spider, error)
	AttackSpider(damage int) error
}

func NewSpiderRepository() spiderRepo {
	return spiderRepo{}
}

func (r spiderRepo) GetSpiders() ([]model.Spider, error) {
	return nil, nil
}

func (r spiderRepo) GetSpiderByUUID(uuid uuid.UUID) (model.Spider, error) {
	return model.Spider{}, nil
}

func (r spiderRepo) AttackSpider(damage int) error {
	return nil
}
