package exercise_two_a

import "fmt"

func Execute() {

	fmt.Println("exercise two a")

	input := make(chan int)
	output := make(chan int, 3)

	multiplier := NewMultiplier()

	multiplier.Start(input, output)

	input <- 1
	input <- 2
	input <- 3

	multiplier.Stop()

	fmt.Println(<-output) // 2
	fmt.Println(<-output) // 4
	fmt.Println(<-output) // 6

	close(input)
	close(output)
}
