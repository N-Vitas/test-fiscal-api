package FiscalTestApi

import (
	"fmt"
	"os"
)

type Rules struct {
	Rule  string
	Data  interface{}
	Valid bool
}

func (s *App) Run() {
	for {
		select {
		case receipt := <-s.receipt:
			s.SendRpc(receipt)
		case responce := <-s.responce:
			//fmt.Println("Server : ", string(responce))
			go s.parceResponce(responce)
			//return

		}
	}
}

func (s *App) parceResponce(r []byte) {
	if s.checkRule(r) {
		rule := s.Next()
		if rule.Valid {
			os.Exit(0)
		}
		s.Done(rule)
		return
	}
	fmt.Println("Тест провален")
	s.CloseApp()
}

func (s *App) LoadRuler() {
	s.addRules("LOGIN", NewLogin())
	s.addRules("ERR_SECRET", NewErrorSecret())
	s.addRules("ERR_VERSION", NewErrorVersion())
	s.addRules("ERR_HEADER", NewErrorHeader())
	s.addRules("VERSION", NewVersion())
	s.addRules("SECTION", NewSection())
	s.addRules("TERMINAL", NewTerminal())
	s.addRules("PRINT", NewPrint())
	s.addRules("CASHIER", NewCashier())
	s.addRules("XREPORT", NewReport())
	s.addRules("CASH_IN", NewCashIn())
	s.addRules("CASH_OUT", NewCashOut())
	s.addRules("XREPORT", NewReport())
	s.addRules("PAYMENT", s.NewPayment())
	s.addRules("BUY", s.NewBuy())
	s.addRules("COMMIT_PAYMENT", s.NewCommitPayment())
	s.addRules("COMMIT_BUY", s.NewCommitBuy())
	s.addRules("CLOSEDAY", NewCloseDay())
}

func (s *App) addRules(name string, data interface{}) {
	s.rules = append(s.rules, Rules{name, data, false})
}

func (s *App) Next() Rules {
	for k, rules := range s.rules {
		if rules.Rule == "COMMIT_BUY" {
			buy := s.NewCommitBuy()
			s.rules[k].Data = buy
			rules.Data = buy
		}
		if rules.Rule == "COMMIT_PAYMENT" {
			payment := s.NewCommitPayment()
			s.rules[k].Data = payment
			rules.Data = payment
		}
		if rules.Valid == false {
			fmt.Println("Тест ", rules.Rule)
			s.receipt <- rules.Data
			return rules
		}
	}
	return Rules{"stop", nil, true}
}

func (s *App) Done(rules Rules) {
	for key, rule := range s.rules {
		if rule.Rule == rules.Rule {
			s.rules[key].Valid = true
		}
	}
}
