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
			fmt.Println("Server : ", string(responce))
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
	os.Exit(1)
}

func (s *App) LoadRuler() {
	s.addRules("login", NewLogin())
	s.addRules("errSecret", NewErrorSecret())
	s.addRules("errVersion", NewErrorVersion())
	s.addRules("errHeader", NewErrorHeader())
	s.addRules("version", NewVersion())
	s.addRules("section", NewSection())
	s.addRules("terminal", NewTerminal())
	s.addRules("print", NewPrint())
	s.addRules("casheir", NewCashier())
	s.addRules("x-report", NewReport())
	s.addRules("payments", s.NewPayment())
	s.addRules("buy", s.NewBuy())
	s.addRules("commitPayment", s.NewCommitPayment())
	s.addRules("commitBuy", s.NewCommitBuy())
}

func (s *App) addRules(name string, data interface{}) {
	s.rules = append(s.rules, Rules{name, data, false})
}

func (s *App) Next() Rules {
	for k, rules := range s.rules {
		if rules.Rule == "commitBuy" {
			buy := s.NewCommitBuy()
			fmt.Println("BUY DATA", buy)
			s.rules[k].Data = buy
			rules.Data = buy
		}
		if rules.Rule == "commitPayment" {
			payment := s.NewCommitPayment()
			s.rules[k].Data = payment
			rules.Data = payment
		}
		if rules.Valid == false {
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
