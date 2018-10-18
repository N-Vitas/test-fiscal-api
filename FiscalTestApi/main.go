package FiscalTestApi

import (
	"net"
	"log"
	"os"
	"encoding/json"
	"fmt"
)

type App struct {
	client net.Conn
	Payments []PaymentRequest `json:"payments"`
	Ch chan map[int]int
	SendJs chan map[string][]byte
	CountCh chan int
	CountRequest int
}

func NewApp() *App {
	return &App{
		nil,
		[]PaymentRequest{},
		make(chan map[int]int),
		make(chan map[string][]byte, 1),
		make(chan int),
		0,
	}
}

func (s *App) ConnectRPC() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:4000")
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	s.client, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("ConnectRPC Error", err)
	}
}

func (s *App) Prepare(col int) {
	for i := 0; i < col; i++ {
		js, err := json.Marshal(s.generateRPCPayment("Виталий"))
		fmt.Println(string(js))
		if err != nil {
			fmt.Println("Ошибка парсинга Prepare",err)
			continue
		}
		s.Send(js)
	}
	// "{\"payments\":[{\"account\": \"3713443856\",\"idService\": 4064,\"amount\": 20190,\"addings\": [ {\"subservice\": 1, \"constraint\": null,\"amount0\": 0,\"amount1\": 0,\"amount2\": 0,\"comission\": 20190,\"memo\": \"MzcxMzQ0Mzg1Ng==\"}],\"amountTare\": 0,\"amountCard\": 0,\"amountCredit\": 0}]}"
}

func (s *App) SendRpc(obj interface{}) {
		js, err := json.Marshal(obj)
		fmt.Println(string(js))
		if err != nil {
			fmt.Println("Ошибка парсинга Prepare",err)
			return
		}
		s.Send(js)
}

func (s *App) PrepareApi(col int,token string) {
	for i := 0; i < col; i++ {
		//fmt.Println(i,token)
		js := s.grouve(i,token)
		if js != nil {
			go s.SendApi(js, token)
			//s.SendJs <- map[string][]byte{token:js}
		}
	}
}
func (s *App) transport(ch map[int]int) {
	s.Ch <- ch
}
type A struct {
	Payments []PaymentRequest `json:"payments"`
}
func (s *App) grouve(col int,token string) []byte {
	a := []PaymentRequest{}
	amount := s.generateAmount(col, token)
	a = append(a,s.generatePayment(amount, token))
	js, err := json.Marshal(A{a})
	if err != nil {
		fmt.Println("Ошибка парсинга Prepare",err)
		return nil
	}
	return js
}
func (s *App) generateAmount(col int,token string) float64 {
	switch token {
	case IRINA:
		return float64(10+col)
	case TERMINAL79320:
		return float64(1+col)
	case TERMINAL79374:
		return float64(500+col)
	case TERMINAL79392:
		return float64(1000+col)
	case TERMINAL79401:
		return float64(1500+col)
	case TERMINAL80064:
		return float64(2000+col)
	}
	return float64(100)
}
func (s *App) Send(resp []byte)  {
	_, err := s.client.Write(resp)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
}
func (s *App) Reader() {
	for {
		reply := make([]byte, 1024)

		_, err := s.client.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			s.client.Close()
			os.Exit(1)
		}
		fmt.Println("Сервер ответил=", string(reply))
	}
}