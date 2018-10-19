package FiscalTestApi

func (s *App) NewCommitBuy() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			OpCode int64
			SysNum int64
		}{5, s.rememer.buyStruct.Status.Transaction.SysNum},
	}
}

type CommitBuy struct {
	Header HRpc
	Status StatusCommitBuy
}
type StatusCommitBuy struct {
	Code        int64
	SysInfo     SysInfo
	Transaction TransactionCommitBuy
}
type TransactionCommitBuy struct {
	Amount   float64
	DateTime string
	SysNum   int64
	Type     int64
}
