package bills

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
this api will take some input
and add a bill in event with given id
*/

type NewBillRequest struct {
	EventId int32 `json:"event_id"`
}

func (s *BillService) NewBill(c echo.Context) error {
	var req NewBillRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "not able to bind the request")
	}

	// to create a new bill and add it to the event
	// sql query to insert the data here

	// return the response 1 for success -1 for error
	return c.JSON(200, "1")
}
