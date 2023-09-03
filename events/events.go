package events

import (
	"reflect"

	"github.com/google/uuid"
)

var Topics = []string{
	reflect.TypeOf(AttackSpider{}).Name(),
}

type Event interface {
}

type AttackSpider struct {
	UUID uuid.UUID
}
