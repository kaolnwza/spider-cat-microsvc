package service

import (
	"github.com/google/uuid"
	"github.com/kaolnwza/spider-cat-microsvc/cat-service/publisher"
)

type spiderSvc struct {
	spiderProd publisher.SpiderProducer
}

type SpiderService interface {
	AttackSpiderService(spiderUUID uuid.UUID) error
}

func NewSpiderService(spiderProd publisher.SpiderProducer) spiderSvc {
	return spiderSvc{spiderProd: spiderProd}
}

func (s spiderSvc) AttackSpiderService(spiderUUID uuid.UUID) error {
	if err := s.spiderProd.AttackSpider(spiderUUID); err != nil {
		return err
	}

	return nil
}
