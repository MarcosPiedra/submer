package exercise_two_c

import (
	"fmt"
	"sync"
	"time"
)

type Multiplier interface {
	Start(input, output chan int)
	Stop()
}

type MultiplierRunner struct {
	chControl chan bool
	chInput   chan int
	chOutput  chan int
	wg        sync.WaitGroup
}

func NewMultiplier() *MultiplierRunner {
	return &MultiplierRunner{}
}

func (s *MultiplierRunner) run() {
	for {
		select {
		case _, ok := <-s.chControl:
			if ok {
				continue
			}
			s.wg.Done()
			return
		case value := <-s.chInput:
			select {
			case s.chOutput <- value * 2:
				break
			case <-time.After(time.Millisecond * 500):
				s.wg.Done()
				return
			}
		}
	}
}

// Start implements Multiplier.
func (s *MultiplierRunner) Start(input chan int, output chan int) {
	s.chControl = make(chan bool)
	s.chInput = input
	s.chOutput = output

	s.wg.Add(1)

	go s.run()
}

// Stop implements Multiplier.
func (s *MultiplierRunner) Stop() {
	close(s.chControl)
	if s.wgTimeout() {
		fmt.Print("Timeout in Stop!")
	}
}

func (s *MultiplierRunner) wgTimeout() bool {
	chCtrl := make(chan bool, 1)

	go func() {
		defer close(chCtrl)
		s.wg.Wait()
	}()

	select {
	case <-chCtrl:
		return false
	case <-time.After(time.Millisecond * 100):
		return true
	}
}

var _ Multiplier = (*MultiplierRunner)(nil)
