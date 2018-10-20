package FiscalTestApi

func NewCashOut() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			Amount float64
			OpCode int64
		}{100, 2},
	}
}

type CashOut struct {
	Header HRpc
	Status StatusCashOut
}
type StatusCashOut struct {
	Code        int64
	SysInfo     SysInfo
	Transaction TransactionCashOut
}
type TransactionCashOut struct {
	Amount   float64
	DateTime string
	SysNum   int64
	Type     int64
}
