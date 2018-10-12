package FiscalTestApi

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"database/sql"
	"time"
	"math/rand"
	"bytes"
)

const IRINA  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjoxOTksIkxvZ2luIjoiZGVtbyIsIlVzZXJUeXBlIjoyLCJGdWxsTmFtZSI6ItCY0YDQuNC90LAiLCJQYXNzd29yZCI6IiIsIlRva2VuIjoiIn0sIkd1aWQiOiJiYzVlMWY4ZjE3NGY1MWYzMzBkMjQwODhiN2ZmMmYwNCIsIklEQWdlbnRzIjozMTQsIklEU3lzVXNlciI6OTQzNTcsIklEVHlwZVRlcm1pbmFsIjoyLCJJZFRlcm1pbmFsIjo3NjYyMCwiU2lnbiI6IiIsImNyZWF0ZWQiOjE1Mzg2MzI4NjF9.V3lkYFURkm_X-1rj0wy5RS3kVsF2F8ZRf1Ya8zMLdDE"
const VITALIY  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjoxOCwiTG9naW4iOiJkZW1vIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiRGVtbyIsIlBhc3N3b3JkIjoiIiwiVG9rZW4iOiIifSwiR3VpZCI6ImU5MmI4MWUyZDg5M2VmOTM5ZGFiZmQ3NDYxZTFmNGNjIiwiSURBZ2VudHMiOjMxNCwiSURTeXNVc2VyIjoxNDY3LCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6MTA2OSwiU2lnbiI6IiIsImNyZWF0ZWQiOjE1MzkxNjgzODN9.71D0If9ZihwioM5b-zprozo75hPWP_A0p4FvQuvMsMA"
const TERMINAL79320 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjo0MTEsIkxvZ2luIjoiY2FzaGVyIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiY2FzaGVyIiwiUGFzc3dvcmQiOiIiLCJUb2tlbiI6IiJ9LCJHdWlkIjoiOTA4NjIzMzhmZTUwZjY3YWQzNjc3MmVjODI4MzU5YTUiLCJJREFnZW50cyI6MzE0LCJJRFN5c1VzZXIiOjk4MDY2LCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6NzkzMjAsIlNpZ24iOiIiLCJjcmVhdGVkIjoxNTM5MzE1OTc4fQ.vfcx9VsUWKYN2xq3kCHULmueXiw_OFN9oTdFpBSgo9M"
const TERMINAL79374 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjo0MTIsIkxvZ2luIjoiY2FzaGVyIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiY2FzaGVyIiwiUGFzc3dvcmQiOiIiLCJUb2tlbiI6IiJ9LCJHdWlkIjoiMDIwOTM1YmE0YjdiOTZjZjBlM2EyNWI3Yjc0Zjk5YTIiLCJJREFnZW50cyI6MzE0LCJJRFN5c1VzZXIiOjk4MTQwLCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6NzkzNzQsIlNpZ24iOiIiLCJjcmVhdGVkIjoxNTM5MzE1ODYyfQ.K3L4hWQvoXoJ0CqlajBYGDL41z0cAUzVFqgvIqxwIYE"
const TERMINAL79392 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjo0MTMsIkxvZ2luIjoiY2FzaGVyIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiY2FzaGVyIiwiUGFzc3dvcmQiOiIiLCJUb2tlbiI6IiJ9LCJHdWlkIjoiNWFkMDQxNWMzZGYwNmNiZGFhNmU2ZjljOGViMTU5NTQiLCJJREFnZW50cyI6MzE0LCJJRFN5c1VzZXIiOjk4MTYwLCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6NzkzOTIsIlNpZ24iOiIiLCJjcmVhdGVkIjoxNTM5MzE1OTAwfQ.lLO5C4jvLu3TOzZmYK4hXCliJIvpu3yGFO0w8465iG4"
const TERMINAL79401 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjo0MTQsIkxvZ2luIjoiY2FzaGVyIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiY2FzaGVyIiwiUGFzc3dvcmQiOiIiLCJUb2tlbiI6IiJ9LCJHdWlkIjoiZDg1M2VhYTUzMGYyYTgyNTFjMWU3YzI1OTE0YWQ4N2IiLCJJREFnZW50cyI6MzE0LCJJRFN5c1VzZXIiOjk4MTc0LCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6Nzk0MDEsIlNpZ24iOiIiLCJjcmVhdGVkIjoxNTM5MzE1OTIxfQ.L2QZh30ASKXLP3kbgXn6lc1aDvGh1kp7qVIgkzo5Hts"
const TERMINAL80064 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXNjYWxBdXRoIjp7IklkIjo0MTUsIkxvZ2luIjoiY2FzaGVyIiwiVXNlclR5cGUiOjIsIkZ1bGxOYW1lIjoiY2FzaGVyIiwiUGFzc3dvcmQiOiIiLCJUb2tlbiI6IiJ9LCJHdWlkIjoiYTllZDFjNWY2N2EzNjBiYWVjZmMwYWYzYjQ5NDc2OGEiLCJJREFnZW50cyI6MzE0LCJJRFN5c1VzZXIiOjk5MTY0LCJJRFR5cGVUZXJtaW5hbCI6MiwiSWRUZXJtaW5hbCI6ODAwNjQsIlNpZ24iOiIiLCJjcmVhdGVkIjoxNTM5MzE1OTQzfQ.njXbIk9_VdHycDN9IMolw4VRUIN8QQF72ptL7dtp0BI"
type PaymentRequest struct{
	SN int `json:"numTrans"`
	Account string `json:"account"`
	Amount float64 `json:"amount"`
	AmountTare float64 `json:"amountTare"`
	AmountCard float64 `json:"amountCard"`
	AmountCredit float64 `json:"amountCredit"`
	Date int `json:"date"`
	IDService int `json:"idService"`
	Post int `json:"post"`
	Currency string `json:"currency"`
	PS string `json:"ps"`
	Addings []Adding `json:"addings"`
}
type Adding struct {
	Memo Memo `json:"memo"`
}
type Memo struct {
	PayData PayData `json:"payData" xml:"payData"`
}
type PayData struct {
	OperCode string `json:"operCode" xml:"operCode,attr"`
	Domain int64 `json:"domain" xml:"domain,attr"`
	Taken float64 `json:"taken" xml:"taken,attr"` // Сумма внесения
	Change float64 `json:"change" xml:"change,attr"` // Сумма сдачи
	PayTypes PayTypes `json:"payTypes" xml:"payTypes"` // Как поатили нал банк кредит и тд.
	Positions Positions `json:"poses" xml:"poses"` // Позиции в чеке. Детализация
	IsOffline bool `json:"isOffline" xml:"isOffline,attr"`
}
type PayTypes struct {
	PayType []PayType `json:"payType" xml:"payType"`
}
type Positions struct {
	Position []Position `json:"pos" xml:"pos"`
}
type PayType struct {
	Type int64 `json:"type" xml:"type,attr"`
	Amount float64 `json:"amount" xml:"amount,attr"`
}
type Position struct {
	Storno bool `json:"isStorno" xml:"isStorno,attr"`
	TaxPercent float64 `json:"percent" xml:"percent,attr"`  // Сам процент НДС от терминала
	Code int64 `json:"posCode" xml:"posCode,attr"`
	Name string `json:"posName" xml:"posName,attr"`
	Section string `json:"posSection" xml:"posSection,attr"`
	Count float64 `json:"posCnt" xml:"posCnt,attr"`
	Price float64 `json:"posPrice" xml:"posPrice,attr"`
	Vat float64 `json:"posNDS" xml:"posNDS,attr"` // Сумма НДС терминала должна калькулироваться
	Discount float64 `json:"posSkidkaSum" xml:"posSkidkaSum,attr"`// Сумма Скидки должна калькулироваться
	Markup float64 `json:"posNacenkaSum" xml:"posNacenkaSum,attr"` // Сумма Наценки надбавка должна калькулироваться
	DiscountVat float64 `json:"posSkidkaNDS" xml:"posSkidkaNDS,attr"` // Сумма НДС от скидки
	MarkupVat float64 `json:"posNacenkaNDS" xml:"posNacenkaNDS,attr"` // Сумма НДС от наценки
}

func (s *App) generatePayment(amount float64, token string) PaymentRequest {
	casher := "casher"
	ts := int(time.Now().Unix())
	sn := ts * (rand.Intn(999999)+1)
	//if token == IRINA {
	//	casher = "Ирина"
	//}
	adings := []Adding{}
	types := []PayType{}
	poses := []Position{}
	//amount := float64(rand.Intn(100))
	types = append(types,PayType{0,amount})
	poses = append(poses,Position{
		false,
		0,
		1,
		"Позиция",
		"section1",
		1,
		amount,
		0,
		0,
		0,
		0,
		0,
	})
	adings = append(adings,Adding{Memo{PayData{
		casher,
		0,
		amount,
		0,
		PayTypes{types},
		Positions{poses},
		true,
	}}})
	//fmt.Printf("Отправлена сумма %v транзакция %d из %d\n",amount, sn, len(s.Count))
	return PaymentRequest{
		Account:"Продажа",
		Amount:amount,
		SN:sn,
		AmountTare:0,
		AmountCard:0,
		AmountCredit:0,
		IDService:167,
		Post:0,
		Currency:"KZT",
		PS:"R",
		Addings:adings,
	}
}
type PaymentStatus struct {
	SN int
	IDTerminal int
	Status int
	DateOut string
	DateIn string
	Finality int
	FiscalAttribute sql.NullString
}
func (s *App) SendApi(js []byte, token string) {
	s.CountCh <- 1
	url := "https://fiskal.api.kassa24.kz/payments/service"
	rm := []PaymentStatus{}
	m := make(map[int]int)
	m[1] = 1
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		//fmt.Println("Ошибка NewRequest: ", err)
		s.transport(m)
		return
	}
	req.Header.Set("authorization", token)
	req.Header.Set("Content-Type", "application/json")
	//http.DefaultClient.Timeout = 2 * time.Minute
	timeout := time.Duration(60 * time.Second)
	client := http.Client{Timeout:timeout}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Ошибка ответа: ", err.Error())
		s.transport(m)
		return
	}
	//defer res.Body.Close()
	//for{
	decoder, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println("Ошибка чтения ответа:", err.Error())
		s.transport(m)
		return
	}
	//fmt.Printf("Ответ %d : %d\n",string(decoder))
	// Парсим декодированный запрос
	err = json.Unmarshal(decoder, &rm)
	if err != nil {
		fmt.Println("Ошибка распарсить josn:", err.Error())
		s.transport(m)
		return
	}
	// проверяем ответ
	r := rm[0]
	if r != (PaymentStatus{}) {
		//fmt.Printf("Транзакция %d успешна\n", r.SN)
		m = make(map[int]int)
		m[r.SN] = r.IDTerminal
		s.transport(m)
		return
	} else {
		s.transport(m)
	}
	//}
}