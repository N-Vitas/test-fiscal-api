package FiscalTestApi

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type HeadRPC struct {
	Header    HRpc
	Operation RPC
}
type HeadDefRPC struct {
	Header    HRpc
	Operation interface{}
}

type HeadRPCInfo struct {
	Header  HRpc
	InfoReq InfoReq
}
type HRpc struct {
	Version int64
	Secret  string
}
type InfoReq struct {
	ReqCode int64
	SysNum  int64
	Receipt string
}

type RPC struct {
	OpCode       int64
	AmountCash   float64
	AmountTare   float64
	AmountCard   float64
	AmountCredit float64
	Change       float64
	Article      []Article
}

type Article struct {
	IsStorno  bool
	CRSection string
	Name      string
	Count     int64
	Price     float64
	Discount  float64
	Charge    float64
}
type SysInfo struct {
	AgentName string
	IinBin    int64
	NdsCert   string
	Rnm       int64
	TermAddr  string
	TermId    int64
}

func (s *App) generateRPCPayment(casher string) HeadRPC {
	rand.Seed(time.Now().UnixNano())
	amount := float64(rand.Intn(100))
	adings := []Article{}
	adings = append(adings, Article{
		false,
		"section1",
		"Продажа",
		1,
		amount,
		0,
		0,
	})
	return HeadRPC{
		Header:    HRpc{1, "1Q2w3e4r"},
		Operation: RPC{3, amount, 0, 0, 0, 0, adings}}
}
func (s *App) generateArticle(i int) (float64, Article) {
	amount := float64(rand.Intn(100))
	d := float64(rand.Intn(10))
	n := float64(rand.Intn(15))
	c := int64(rand.Intn(5) + 1)
	ading := Article{
		false,
		"section2",
		fmt.Sprintf("Позиция %d", i),
		c,
		amount,
		d,
		n,
	}
	// Сумма позиции
	aa := amount * float64(c)
	// Скидка
	dd := (aa * d) / 100
	// наценка
	nn := (aa * n) / 100
	// Сумма + наценка - скидка
	cash := (aa + nn) - dd
	return cash, ading
}
func (s *App) ceil(cash float64) float64 {
	n := fmt.Sprintf("%.2f", cash)
	cash, _ = strconv.ParseFloat(n, 2)
	return cash
}

func (d *App) round(x float64) float64 {
	const (
		mask  = 0x7FF
		shift = 64 - 11 - 1
		bias  = 1023

		signMask = 1 << 63
		fracMask = (1 << shift) - 1
		halfMask = 1 << (shift - 1)
		one      = bias << shift
	)

	bits := math.Float64bits(x)
	e := uint(bits>>shift) & mask
	switch {
	case e < bias:
		// Round abs(x)<1 including denormals.
		bits &= signMask // +-0
		if e == bias-1 {
			bits |= one // +-1
		}
	case e < bias+shift:
		// Round any abs(x)>=1 containing a fractional component [0,1).
		e -= bias
		bits += halfMask >> e
		bits &^= fracMask >> e
	}
	return math.Float64frombits(bits)
}

func (s *App) generateRPCPaymentNds() HeadRPC {
	rand.Seed(time.Now().UnixNano())
	cash := float64(0)
	card := float64(0)
	cred := float64(0)
	tara := float64(0)
	adings := []Article{}
	for i := 1; i <= 4; i++ {
		amount, ading := s.generateArticle(i)
		adings = append(adings, ading)
		cash = cash + amount
	}
	return HeadRPC{
		Header:    HRpc{1, "1Q2w3e4r"},
		Operation: RPC{3, s.round(cash), tara, card, cred, 0, adings}}
}
func (s *App) generateRPCPaymentNdsCard() HeadRPC {
	rand.Seed(time.Now().UnixNano())
	cash := float64(0)
	card := float64(0)
	cred := float64(0)
	tara := float64(0)
	adings := []Article{}
	for i := 1; i <= 5; i++ {
		amount, ading := s.generateArticle(i)
		adings = append(adings, ading)
		if i < 3 {
			cash = cash + amount
		}
		card = card + amount
	}
	return HeadRPC{
		Header:    HRpc{1, "1Q2w3e4r"},
		Operation: RPC{3, s.round(cash), tara, s.ceil(card - s.round(cash)), cred, 0, adings}}
}
