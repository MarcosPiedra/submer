package exercise_two_a

import (
	"sync"
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
		case <-s.chControl:
			s.wg.Done()
			return
		case value := <-s.chInput:
			s.chOutput <- value * 2
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
	s.wg.Wait()
}

var _ Multiplier = (*MultiplierRunner)(nil)
