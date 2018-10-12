package Test1

import "fmt"

type Test1 struct {
	maps int
	stop int
	ch chan string
	rune chan int
	done chan bool
}
func NewTest1() *Test1 {
	return &Test1{0,0,make(chan string),make(chan int),make(chan bool)}
}

func (s *Test1) Start(maps int) {
	s.maps = maps
	for i := 0; i < maps ; i++  {
		go s.Add(i)
	}
}
func (s *Test1) Listner() {
	for {
		select {
			case msg := <-s.ch:
				fmt.Println("Пришел", msg)
			break
			case r := <-s.rune:
				go s.Transport(fmt.Sprintf("Транспорт %d", r))
				s.Stop()
			break
			case d := <-s.done:
				if d {
					close(s.ch)
					close(s.rune)
					close(s.done)
					return
				}
			break
		}
	}
}

func (s *Test1) Add(num int) {
	s.rune <- num
}

func (s *Test1) Transport(msg string) {
	s.ch <- msg
}

func (s *Test1) Stop() {
	s.stop++
	if s.maps == s.stop {
		s.done <- true
	}
}