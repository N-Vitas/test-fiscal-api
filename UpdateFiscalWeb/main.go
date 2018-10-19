package UpdateFiscalWeb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const URL_WEB_APP = "http://pult24.vitas/login?return=%2Fdashboard"

type Context struct {
	findStr string
	succes  int
	danger  int
	res     chan bool
	count   chan bool
}

func New() Context {
	return Context{"<div class=\"d-flex flex-column flex\">", 0, 0, make(chan bool), make(chan bool)}
}

func (s *Context) Start(count int) {
	fmt.Println("Start", count)
	for i := 0; i < count; i++ {
		go s.getContext()
	}
	m := 0
	r := 0
	for {
		select {
		case success := <-s.res:
			if success {
				s.succes++
			} else {
				s.danger++
			}
			m++
		case <-s.count:
			r++
			fmt.Printf("Отправка запроса %d \r", r)
		}
		if m >= count {
			close(s.res)
			fmt.Println()
			return
		}
	}
}

func (s *Context) getContext() {
	s.goRequerst()
	req, err := http.NewRequest("GET", URL_WEB_APP, bytes.NewBuffer(nil))
	//http.DefaultClient.Timeout = 2 * time.Minute
	timeout := time.Duration(10 * time.Second)
	client := http.Client{Timeout: timeout}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Ошибка ответа: %s \r", err.Error())
		s.goResponce(false)
		return
	}
	//defer res.Body.Close()
	//for{
	decoder, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %s \r", err.Error())
		s.goResponce(false)
		return
	}
	if strings.Index(string(decoder), s.findStr) != -1 {
		s.goResponce(true)
	} else {
		s.goResponce(false)
	}
}

func (s *Context) goRequerst()           { s.count <- true }
func (s *Context) goResponce(valid bool) { s.res <- valid }

func (s *Context) GetSuccess() int { return s.succes }
func (s *Context) GetError() int   { return s.danger }
