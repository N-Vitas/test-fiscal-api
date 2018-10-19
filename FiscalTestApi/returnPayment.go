package FiscalTestApi

func (s *App) NewCommitPayment() HeadDefRPC {
	return HeadDefRPC{
		HRpc{1, "1Q2w3e4r"},
		struct {
			OpCode int64
			SysNum int64
		}{6, s.rememer.paymentStruct.Status.Transaction.SysNum},
	}
}

type CommitPayment struct {
	Header HRpc
	Status StatusCommitPayment
}
type StatusCommitPayment struct {
	Code        int64
	SysInfo     SysInfo
	Transaction TransactionCommitPayment
}
type TransactionCommitPayment struct {
	Amount   float64
	DateTime string
	SysNum   int64
	Type     int64
}
