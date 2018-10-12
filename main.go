package main

import (
	"test/FiscalTestApi"
	"fmt"
)

type Epocha struct {
	ch int
	count int
}

func (s *Epocha) GetChanelCount() int {
	return s.count/s.ch
}

func (s *Epocha) GetCount() int {
	return s.count
}
func main()  {
	epoha := Epocha{2,10000}
	finish := 0
	err := 0
	s := FiscalTestApi.NewApp()
	//s.ConnectRPC()
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.VITALIY)
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.IRINA)
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79320)
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79374)
	//go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79392)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL79401)
	go s.PrepareApi(epoha.GetChanelCount(),FiscalTestApi.TERMINAL80064)
	//defer close(s.Ch)
	//s.Prepare(100)
	//s.Reader()
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
		if finish + err == epoha.GetCount() {
			return
		}
	}
}