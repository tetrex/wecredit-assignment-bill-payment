package bills

import "github.com/labstack/echo/v4"

/*
this is get the bill split for each user

ex:
bill = 100
users = 3

response = [

	{ user_id: 1, amount: 33.33 },
	{ user_id: 2, amount: 33.33 },
	{ user_id: 3, amount: 33.33 },

]
*/
func (s *BillService) GetBillSplit(c echo.Context) error {
	/*
		bill by bill id
		get total users in the bill
		just divide the bill amount by total users
		return the response
	*/
	return nil
}
