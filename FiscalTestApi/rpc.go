package FiscalTestApi

import "math/rand"

type HeadRPC struct {
	Header HRpc
	Operation RPC
}

type HRpc struct {
	Version int64
	Secret string
}
type RPC struct {
	OpCode int64
	AmountCash float64
	AmountTare float64
	AmountCard float64
	AmountCredit float64
	Change float64
	Article []Article
}

type Article struct {
	IsStorno bool
	CRSection string
	Name string
	Count int64
	Price float64
	Discount float64
	Charge float64
}


func (s *App) generateRPCPayment(casher string) HeadRPC {
	amount := float64(rand.Intn(100))
	adings := []Article{}
	adings = append(adings,Article{
		false,
		"section1",
		"Продажа",
		1,
		amount,
		0,
		0,
	})
	return HeadRPC{
		Header:HRpc{1,"1Q2w3e4r"},
		Operation:RPC{3,amount,0,0,0,0,adings}}
}

