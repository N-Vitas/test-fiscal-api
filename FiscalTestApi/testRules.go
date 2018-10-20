package FiscalTestApi

import (
	"fmt"
	"strings"
	"encoding/json"
)

const (
	START_DEMO     = `{"Header":{"Secret":"","Version":1},"Status":{"Code":0,"Message":{"dateCreate":"06.09.2018 18:02:03","fiscalId":199,"login":"demo","name":"Ирина","status":1,"userType":"cashier"}}}`
	ERR_SECRET     = `{"Header":{"Secret":"","Version":1},"Status":{"Code":1,"Message":"Не верный секретный ключ"}}`
	ERR_VERSION    = `{"Header":{"Secret":"","Version":1},"Status":{"Code":1,"Message":"Не совподают версии"}}`
	ERR_HEADER     = `{"Header":{"Secret":"","Version":1},"Status":{"Code":1,"Message":"Отсутствует заголовок"}}`
	START_NOCASHER = `{"Header":{"Secret":"","Version":1},"Status":{"Code":1,"Message":"Пожалуйста авторизируйтесь в приложении"}}`
	LOGIN          = `{"Header":{"Secret":"","Version":1},"Status":{"Code":0,"Message":{"dateCreate":"06.09.2018 18:02:03","fiscalId":199,"login":"demo","name":"Ирина","status":1}}}`
	VERSION        = `{"Header":{"Secret":"","Version":1},"Status":{"Code":0,"Message":"Fiscal TCP/IP version 1"}}`
	SECTION        = `{"Header":{"Secret":"","Version":1},"Status":{"Code":0,"Sections":[{"key":"section1","title":"Без ндс","value":0},{"key":"section2","title":"Секция 1","value":0},{"key":"section3","title":"Секция 2","value":0},{"key":"section4","title":"Секция 3","value":0},{"key":"section5","title":"Секция 4","value":0},{"key":"section6","title":"Секция 5","value":0}]}}`
	TERMINAL       = `{"Header":{"Secret":"","Version":1},"Status":{"Code":0,"Message":{"AddressCompany":"ул.Чкалова 48, оф.324","AddressPoint":"test address","AddressSupport":"tech address","DataNDS":"082467999","EnableSimRequest":0,"FisCode":1,"IDTerminal":76620,"IdLocation":11,"Iin":150341016439,"IsSystem":0,"Kfu":182,"KsRegister":314,"NameDiler":"TestAgent TOO","NamePoint":"TestTerminal для оплат","Nds":0,"Region":"Петропавловск","Rnm":132465798123,"Rnn":123456789123,"Rnnfil":0,"TaxIDInspection":301,"TelSupport":"tech phone","TerminalAddress":"test address","TerminalName":"TestTerminal для оплат","Version_po":""}}}`
	PRINT          = `Печать отправлена на принтер`
	PAYMENT        = 3
	BUY            = 4
	COMMIT_BUY     = 5
	COMMIT_PAYMENT = 6
	CASH_IN        = 1
	CASH_OUT       = 2
	CLOSEDAY       = `Смена успешно закрыта`
)

func (s *App) checkRule(reply []byte) bool {
	responce := string(reply)
	if strings.Index(responce, START_DEMO) != -1 {
		fmt.Println("START_DEMO OK")
		return true
	}
	if strings.Index(responce, ERR_SECRET) != -1 {
		fmt.Println("ERR_SECRET OK")
		return true
	}
	if strings.Index(responce, ERR_VERSION) != -1 {
		fmt.Println("ERR_VERSION OK")
		return true
	}
	if strings.Index(responce, ERR_HEADER) != -1 {
		fmt.Println("ERR_HEADER OK")
		return true
	}
	if strings.Index(responce, ERR_SECRET) != -1 {
		fmt.Println("ERR_SECRET OK")
		return true
	}
	if strings.Index(responce, START_NOCASHER) != -1 {
		fmt.Println("START_NOCASHER OK")
		return true
	}
	if strings.Index(responce, LOGIN) != -1 {
		fmt.Println("LOGIN/CASHIER OK")
		return true
	}
	if strings.Index(responce, VERSION) != -1 {
		fmt.Println("VERSION OK")
		return true
	}
	if strings.Index(responce, SECTION) != -1 {
		fmt.Println("SECTION OK")
		return true
	}
	if strings.Index(responce, TERMINAL) != -1 {
		fmt.Println("TERMINAL OK")
		return true
	}
	if strings.Index(responce, PRINT) != -1 {
		fmt.Println("PRINT OK")
		return true
	}
	if strings.Index(responce, CLOSEDAY) != -1 {
		fmt.Println("CLOSEDAY OK")
		return true
	}
	if s.rememer.CheckReport(reply) {
		fmt.Println("XREPORT OK")
		return true
	}
	if s.getOperation(reply, PAYMENT) {
		if s.rememer.CheckPayment(reply) {
			fmt.Println("PAYMENTS OK")
			return true
		}
	}
	if s.getOperation(reply, BUY) {
		if s.rememer.CheckBuy(reply) {
			fmt.Println("BUY OK")
			return true
		}
	}
	if s.getOperation(reply, COMMIT_PAYMENT) {
		if s.rememer.CheckCommitPayment(reply) {
			fmt.Println("COMMIT_PAYMENT OK")
			return true
		}
	}
	if s.getOperation(reply, COMMIT_BUY) {
		if s.rememer.CheckCommitBuy(reply) {
			fmt.Println("COMMIT_BUY OK")
			return true
		}
	}
	if s.getOperation(reply, CASH_IN) {
		if s.rememer.CheckCashIn(reply) {
			fmt.Println("CASH_IN OK")
			return true
		}
	}
	if s.getOperation(reply, CASH_OUT) {
		if s.rememer.CheckCashOut(reply) {
			fmt.Println("CASH_OUT OK")
			return true
		}
	}
	return false
}

func (s *App) getOperation(report []byte, operation int64) bool {
	o := Operation{}
	if err := json.Unmarshal(report, &o); err != nil {
		fmt.Println(err)
		return false
	}
	return operation == o.Status.Transaction.Type
}