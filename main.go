package main

import (
	//"test-fiscal-api/UpdateFiscalWeb"
	"test-fiscal-api/FiscalTestApi"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

type Epocha struct {
	ch    int
	count int
}

func (s *Epocha) GetChanelCount() int {
	return s.count / s.ch
}

func (s *Epocha) GetCount() int {
	return s.count
}
func cpuhogger() {
	var acc uint64
	for {
		acc += 1
		if acc&1 == 0 {
			acc <<= 1
		}
	}
}

func main() {
	go http.ListenAndServe("0.0.0.0:8082", nil)
	cpuhogger()
}
func mainDeb() {
	epoha := Epocha{4,8000}
	//test := UpdateFiscalWeb.New()
	//test.Start(epoha.GetCount())
	//fmt.Printf("Результат %d из %d успешны\n",epoha.GetChanelCount(), 40)
	s := FiscalTestApi.NewApp()
	defer s.CloseApp()
	//fmt.Printf("Результат %d из %d ошибок\n",test.GetError(),epoha.GetCount())
	finish := 0
	err := 0
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.VITALIY)
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.IRINA)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79320)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79392)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79401)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL80064)
	for {
		select {
		case maps := <-s.Ch:
			for sn, term := range maps {
				if term != 1 {
					finish++
					fmt.Printf("Результат %d из %d. SN %d\n", finish,epoha.GetCount(),sn)
				} else {
					err++
					fmt.Printf("Ошибки %d из %d\n", err,epoha.GetCount())
				}
			}
		case maps2 := <-s.SendJs:
			for token, body := range maps2 {
				go s.SendApi(body, token)
			}

		case <-s.CountCh:
			s.CountRequest++
			fmt.Println("Отправка запроса ",s.CountRequest)
		}
		if finish + err >= epoha.GetCount() {
			return
		}
	}

	//s.LoadRuler()
	////rule := s.Next()
	////fmt.Println(rule)
	////s.Done(rule)
	////rule = s.Next()
	////fmt.Println(rule)
	////s.Done(rule)
	////rule = s.Next()
	////fmt.Println(rule)
	//s.ConnectRPC()
	//defer close(s.Ch)
	////s.Prepare(epoha.GetCount())
	//go s.Reader()
	//s.Run()
}
