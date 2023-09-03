package repository

import "github.com/kaolnwza/spider-cat-microsvc/cat-service/model"

type catRepo struct{}

type CatRepository interface {
	GetHP() (model.Cat, error)
}

func NewCatRepository() catRepo {
	return catRepo{}
}

func (r catRepo) GetHP() (model.Cat, error) {
	return model.Cat{HP: 100}, nil
}
