package FiscalTestApi

import "fmt"

func (s *App) NewPayment() HeadRPC {
	payment := s.generateRPCPayment("Demo")
	s.rememer.Amount = payment.Operation.AmountCash
	return payment
}
func (s *App) NewPaymentNds() HeadRPC {
	payment := s.generateRPCPaymentNds()
	fmt.Printf("NewPaymentNds %v\n", payment)
	s.rememer.Amount = payment.Operation.AmountCash
	return payment
}
func (s *App) NewPaymentNdsCard() HeadRPC {
	payment := s.generateRPCPaymentNdsCard()
	fmt.Printf("NewPaymentNdsCard %v\n", payment)
	s.rememer.Amount = payment.Operation.AmountCash + payment.Operation.AmountCard
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
	Fiscal   int64
	SysNum   int64
	Type     int64
}
