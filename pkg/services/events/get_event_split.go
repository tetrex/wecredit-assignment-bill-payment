package events

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventSplitRequest struct {
	EventId string `json:"event_id"`
}

/*
this api will give us totals for a event
including
1. total amount
2. total users
3. total bills
4. total paid
5. total unpaid
*/
func (s *EventService) EventsSplit(c echo.Context) error {
	var req EventSplitRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "not able to bind the request")
	}

	/*
		step 1. get the event details
		step 2. get all bills by event id
		step 3. loop over all bills and get the total amount and total paid
		step 4. just sum the total amount and total paid and calculate unpaid accordingly [ substract the paid amount from total amount , for each bill and total ]

		and send the response

	*/
	// return the response 1 for success -1 for error
	return c.JSON(200, "1")
}
