package services

import (
	"bills_split/pkg/services/events"
)

type RouterParams struct {
	Event_service events.EventService
}

func InitRouter(p *RouterParams) {

}
