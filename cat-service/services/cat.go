package service

import (
	"github.com/kaolnwza/spider-cat-microsvc/cat-service/model"
	repository "github.com/kaolnwza/spider-cat-microsvc/cat-service/repositories"
)

type catSvc struct {
	catRepo repository.CatRepository
}

type CatService interface {
	GetCatHPService() (model.Cat, error)
}

func NewCatService(catRepo repository.CatRepository) catSvc {
	return catSvc{catRepo: catRepo}
}

func (s catSvc) GetCatHPService() (model.Cat, error) {
	return s.catRepo.GetHP()
}
