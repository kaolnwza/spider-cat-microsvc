package subscriber

import (
	"encoding/json"
	"events"
	"log"
	"reflect"

	service "github.com/kaolnwza/spider-cat-microsvc/spider-service/services"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type spiderEventHandler struct {
	spiderSvc service.SpiderService
}

func NewSpiderEventHandler(spiderSvc service.SpiderService) EventHandler {
	return spiderEventHandler{spiderSvc: spiderSvc}
}

func (obj spiderEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.AttackSpider{}).Name():
		event := &events.AttackSpider{}
		if err := json.Unmarshal(eventBytes, event); err != nil {
			log.Println(err)
			return
		}

		if err := obj.spiderSvc.AttackSpiderService(event.UUID, 555); err != nil {
			log.Println(err)
			return
		}

		log.Printf("[%v] %#v", topic, event)
	default:
		log.Println("no event handler")
	}
}
