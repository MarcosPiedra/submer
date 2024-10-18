package exercise_two_b

import (
	"fmt"
	"time"
)

func Execute() {
	fmt.Println("exercise two b")

	input := make(chan int)
	output := make(chan int, 3)

	multiplier := NewMultiplier()

	multiplier.Start(input, output)

	input <- 1
	input <- 2
	input <- 3
	input <- 4
	time.Sleep(1 * time.Second)

	multiplier.Stop()

	fmt.Println(<-output) // 2
	fmt.Println(<-output) // 4
	fmt.Println(<-output) // 6

	close(input)
	close(output)
}
