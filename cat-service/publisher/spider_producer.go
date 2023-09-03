package publisher

import (
	"events"
	"log"

	"github.com/google/uuid"
)

type SpiderProducer interface {
	AttackSpider(spiderUUID uuid.UUID) error
}

type spiderProd struct {
	eventProducer EventProducer
}

func NewSpiderProducer(eventProducer EventProducer) spiderProd {
	return spiderProd{eventProducer: eventProducer}
}

func (obj spiderProd) AttackSpider(spiderUUID uuid.UUID) error {
	event := events.AttackSpider{}
	event.UUID = spiderUUID

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}
