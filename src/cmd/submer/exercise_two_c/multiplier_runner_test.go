package exercise_two_c

import "testing"

func Test_Multiplier_two_c(t *testing.T) {
	// Arrange
	multiplier := NewMultiplier()
	input := make(chan int)
	output := make(chan int, 3)
	io := [][]int{
		{1, 2},
		{2, 4},
		{3, 6},
	}

	// Act
	multiplier.Start(input, output)

	for _, e := range io {
		input <- e[0]
	}

	multiplier.Stop()

	// Assert
	for _, e := range io {
		o := <-output
		if e[1] != o {
			t.Errorf("expected %v but got %v", e[1], o)
		}
	}

	close(input)
	close(output)
}
