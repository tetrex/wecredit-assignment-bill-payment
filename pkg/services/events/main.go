package events

/*
this will init the service , and reutrn the service instance pointer

input

redis
db_conn
app_env = dev , prod , stage
*/

type EventService struct{}

func NewEventService() EventService {
	return EventService{}
}
