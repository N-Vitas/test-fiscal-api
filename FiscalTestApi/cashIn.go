package FiscalTestApi

func NewCashIn() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			Amount float64
			OpCode int64
		}{100, 1},
	}
}

type CashIn struct {
	Header HRpc
	Status StatusCashIn
}
type StatusCashIn struct {
	Code        int64
	SysInfo     SysInfo
	Transaction TransactionCashIn
}
type TransactionCashIn struct {
	Amount   float64
	DateTime string
	SysNum   int64
	Type     int64
}
