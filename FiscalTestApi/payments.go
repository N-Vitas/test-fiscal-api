package FiscalTestApi

func (s *App) NewPayment() HeadRPC {
	payment := s.generateRPCPayment("Ирина")
	s.rememer.Amount = payment.Operation.AmountCash
	return payment
}

type Payment struct {
	Header HRpc
	Status StatusPayment
}

type StatusPayment struct {
	Code        int64
	SysInfo     SysInfo
	Transaction Transaction
}
type Transaction struct {
	Amount   float64
	DateTime string
	Fiscal   string
	SysNum   int64
	Type     int64
}
