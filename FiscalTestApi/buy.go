package FiscalTestApi

func (s *App) NewBuy() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			OpCode  int64
			Amount  float64
			Account string
		}{4, 100, "Покупка"},
	}
}

type Buy struct {
	Header HRpc
	Status StatusBuy
}

type StatusBuy struct {
	Code        int64
	SysInfo     SysInfo
	Transaction TransactionBuy
}
type TransactionBuy struct {
	DateTime string
	SysNum   int64
	Type     int64
}
