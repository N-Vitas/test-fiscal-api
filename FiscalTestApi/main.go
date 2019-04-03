package FiscalTestApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type App struct {
	client       net.Conn
	Payments     []PaymentRequest `json:"payments"`
	Ch           chan map[int]int
	SendJs       chan map[string][]byte
	CountCh      chan int
	CountRequest int
	receipt      chan interface{}
	responce     chan []byte
	rules        []Rules
	rememer      *Tests
}

func NewApp() *App {
	return &App{
		nil,
		[]PaymentRequest{},
		make(chan map[int]int),
		make(chan map[string][]byte, 1),
		make(chan int),
		0,
		make(chan interface{}),
		make(chan []byte),
		[]Rules{},
		NewRemember(),
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
			fmt.Println("Ошибка парсинга Prepare", err)
			continue
		}
		s.Send(js)
	}
	// "{\"payments\":[{\"account\": \"3713443856\",\"idService\": 4064,\"amount\": 20190,\"addings\": [ {\"subservice\": 1, \"constraint\": null,\"amount0\": 0,\"amount1\": 0,\"amount2\": 0,\"comission\": 20190,\"memo\": \"MzcxMzQ0Mzg1Ng==\"}],\"amountTare\": 0,\"amountCard\": 0,\"amountCredit\": 0}]}"
}

func (s *App) SendRpc(obj interface{}) {
	time.Sleep(time.Second * 5)
	js, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Ошибка парсинга Prepare", err)
		return
	}
	s.Send(js)
}

func (s *App) PrepareApi(col int, token string) {
	for i := 0; i < col; i++ {
		//fmt.Println(i,token)
		js := s.grouve(i, token)
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
	Payments PaymentRequest `json:"payments"`
}

func (s *App) grouve(col int, token string) []byte {
	js, err := json.Marshal(A{s.generatePayment(s.generateAmount(col, token), token)})
	if err != nil {
		fmt.Println("Ошибка парсинга Prepare", err)
		return nil
	}
	return js
}
func (s *App) generateAmount(col int, token string) float64 {
	switch token {
	case IRINA:
		return float64(10 + col)
	case TERMINAL79320:
		return float64(1500 + col)
	case TERMINAL79392:
		return float64(1000 + col)
	case TERMINAL79401:
		return float64(500 + col)
	case TERMINAL80064:
		return float64(10 + col)
	}
	return float64(100)
}
func (s *App) Send(resp []byte) {
	_, err := s.client.Write(resp)
	if err != nil {
		println("Write to server failed:", err.Error())
		s.CloseApp()
		return
	}
}
func (s *App) Reader() {
	for {
		reply := make([]byte, 5000)
		n, err := s.client.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			s.client.Close()
			s.CloseApp()
			return
		}
		s.responce <- reply[:n]
	}
}

func (s *App) CloseApp() {
	fmt.Println("Тест завершен. Завершить работу Y/N")
	clos := "Y"
	fmt.Scan(&clos)
	if strings.ToUpper(clos) == "Y" {
		os.Exit(0)
		return
	}
	s.CloseApp()
}
