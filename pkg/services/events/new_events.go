package events

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
this api will take some input
will crate a new evnet ,
and return some response

// GET
*/

type NewEventsRequest struct {
	EventName string `json:"event_name"`
}

func (s *EventService) NewEvnets(c echo.Context) error {
	var req NewEventsRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "not able to bind the request")
	}

	// to create a new event
	// sql query to insert the data here

	// return the response 1 for success -1 for error
	return c.JSON(200, "1")
}
