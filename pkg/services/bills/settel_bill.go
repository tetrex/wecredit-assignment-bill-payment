package bills

import "github.com/labstack/echo/v4"

/*
this will settel the bill by bill id
request :

	{
		bill_id : 1,
		amounts: [
			{
				user_id : 1,
				amount : 100
			}
		]
	}

amounts contains array of amounts paid by users id

a bill is settel when the user pay the bill , the bill status will be updated to settel auctomaticely
when last payment is done
*/
func (s *BillService) SettelBill(c echo.Context) error {
	/*
		1. get the bill id from the request -> sql function call
		2. add amount to paid coll in table of bill -> if condition
		3. store userid and amounts paid in a table of payments [ for future refrence ] -> sql function call
		4. check if the bill is settel or not [ if settel update the status of bill to settel ] -> if condition and sql call
	*/
	return nil
}
