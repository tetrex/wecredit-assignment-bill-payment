package bills

/*
this will init the service , and reutrn the service instance pointer

input

redis
db_conn
app_env = dev , prod , stage
*/

type BillService struct{}

func NewNewBillService() BillService {
	return BillService{}
}
