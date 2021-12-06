package db

type Borrow struct {
	Address         string
	ContractAddress string
	AccountBorrows  string
	TotalBorrows    string
	Liquility       string
	BlockHeight     int64
	TxHash          string
}
