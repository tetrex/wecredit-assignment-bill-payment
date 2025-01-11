package bills

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
this api will take some input
and remove users in bill with given id
*/

type RemoveUserRequest struct {
	BillId  int32   `json:"bill_id"`
	UserIds []int32 `json:"user_ids"`
}

func (s *BillService) RemoveUser(c echo.Context) error {
	var req RemoveUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "not able to bind the request")
	}

	// to create a new bill and remove it to the event
	// sql query to insert the data here

	// return the response 1 for success -1 for error
	return c.JSON(200, "1")
}
