package FiscalTestApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Tests struct {
	reportMap           []string
	paymentMap          []string
	buyMap              []string
	commitMap           []string
	reportStruct        XReport
	paymentStruct       Payment
	buyStruct           Buy
	commitPaymentStruct CommitPayment
	commitBuyStruct     CommitBuy
	cashInStruct        CashIn
	cashOutStruct       CashOut
	SN                  int64
	Amount              float64
}

func NewRemember() *Tests {
	return &Tests{
		[]string{"Header", "Secret", "Version", "Status", "Code", "Message", "BuyCount", "BuySum", "Com", "HourWork", "InCount", "InSum", "LeftCash", "List", "NDS", "OpenDay", "OutCount", "OutSum", "PayCardSum", "PayCount", "PayCreditSum", "PaySum", "PayTareSum", "ReturnBuyCount", "ReturnBuySum", "ReturnPayCardSum", "ReturnPayCount", "ReturnPayCreditSum", "ReturnPaySum", "ReturnPayTareSum", "SysInfo", "AgentName", "IinBin", "NdsCert", "Rnm", "TermAddr", "TermId"},
		[]string{"Header", "Secret", "Version", "Status", "Code", "SysInfo", "AgentName", "IinBin", "NdsCert", "Rnm", "TermAddr", "TermId", "Transaction", "Amount", "DateTime", "Fiscal", "SysNum", "Type"},
		[]string{"Header", "Secret", "Version", "Status", "Code", "SysInfo", "AgentName", "IinBin", "NdsCert", "Rnm", "TermAddr", "TermId", "Transaction", "DateTime", "SysNum", "Type"},
		[]string{"Header", "Secret", "Version", "Status", "Code", "SysInfo", "AgentName", "IinBin", "NdsCert", "Rnm", "TermAddr", "TermId", "Transaction", "Amount", "DateTime", "SysNum", "Type"},
		XReport{},
		Payment{},
		Buy{},
		CommitPayment{},
		CommitBuy{},
		CashIn{},
		CashOut{},
		0,
		0,
	}
}

func (s *Tests) CheckReport(report []byte) bool {
	return s.checkMap(report, s.reportMap) && s.checkReportData(report)
}

func (s *Tests) CheckPayment(report []byte) bool {
	return s.checkMap(report, s.paymentMap) && s.checkPaymentData(report)
}

func (s *Tests) CheckBuy(report []byte) bool {
	return s.checkMap(report, s.buyMap) && s.checkBuyData(report)
}

func (s *Tests) CheckCommitPayment(report []byte) bool {
	return s.checkMap(report, s.commitMap) && s.checkCommitPaymentData(report)
}

func (s *Tests) CheckCommitBuy(report []byte) bool {
	return s.checkMap(report, s.commitMap) && s.checkCommitBuyData(report)
}

func (s *Tests) CheckCashIn(report []byte) bool {
	return s.checkMap(report, s.commitMap) && s.checkCashIn(report)
}

func (s *Tests) CheckCashOut(report []byte) bool {
	return s.checkMap(report, s.commitMap) && s.checkCashOut(report)
}

func (s *Tests) checkMap(report []byte, arr []string) bool {
	check := 0
	for _, key := range arr {
		if strings.Index(string(report), key) != -1 {
			check++
		}
	}
	return check == len(arr)
}
func (s *Tests) checkReportData(report []byte) bool {
	report = bytes.TrimPrefix(report, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(report, &s.reportStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.reportStruct.Status.Message.LeftCash == 0 {
		return false
	}
	if len(s.reportStruct.Status.Message.OpenDay) == 0 {
		return false
	}
	return true
}
func (s *Tests) checkPaymentData(report []byte) bool {
	if err := json.Unmarshal(report, &s.paymentStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.paymentStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.paymentStruct.Status.Transaction.Amount == 0 {
		fmt.Println("Сумма продажи равна нулю")
		return false
	}
	if s.paymentStruct.Status.Transaction.Type != 3 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.paymentStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд.")
		return true
	}
	return true
}
func (s *Tests) checkBuyData(report []byte) bool {
	if err := json.Unmarshal(report, &s.buyStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.buyStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.buyStruct.Status.Transaction.Type != 4 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.buyStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд")
		return true
	}
	return true
}

func (s *Tests) checkCommitPaymentData(report []byte) bool {
	if err := json.Unmarshal(report, &s.commitPaymentStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.commitPaymentStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.commitPaymentStruct.Status.Transaction.Amount == 0 {
		fmt.Println("Сумма возврата продажи равна нулю")
		return false
	}
	if s.commitPaymentStruct.Status.Transaction.Type != 6 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.commitPaymentStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд.")
		return true
	}
	return true
}
func (s *Tests) checkCommitBuyData(report []byte) bool {
	if err := json.Unmarshal(report, &s.commitBuyStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.commitBuyStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.commitBuyStruct.Status.Transaction.Amount == 0 {
		fmt.Println("Сумма возврата покупки равна нулю")
		return false
	}
	if s.commitBuyStruct.Status.Transaction.Type != 5 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.commitBuyStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд.")
		return true
	}
	return true
}
func (s *Tests) checkCashIn(report []byte) bool {
	if err := json.Unmarshal(report, &s.cashInStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.cashInStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.cashInStruct.Status.Transaction.Amount != 100 {
		fmt.Println("Неверная сумма служебного прихода")
		return false
	}
	if s.cashInStruct.Status.Transaction.Type != 1 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.cashInStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд.")
		return true
	}
	return true
}
func (s *Tests) checkCashOut(report []byte) bool {
	if err := json.Unmarshal(report, &s.cashOutStruct); err != nil {
		fmt.Println(err)
		return false
	}
	if s.cashOutStruct.Status.Transaction.SysNum == 0 {
		fmt.Println("Отсутствует номер транзакции")
		return false
	}
	if s.cashOutStruct.Status.Transaction.Amount != 100 {
		fmt.Println("Неверная сумма служебного расхода")
		return false
	}
	if s.cashOutStruct.Status.Transaction.Type != 2 {
		fmt.Println("Не верный тип транзакции")
		return false
	}
	if len(s.cashOutStruct.Status.Transaction.DateTime) == 0 {
		fmt.Println("Внимание отсутствует дата регистрации в офд.")
		return true
	}
	return true
}
