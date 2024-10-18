package exercise_one

import (
	"fmt"
)

func Execute() {
	fmt.Println("exercise one")

	input := []string{"a", "a", "b", "c", "c", "c", "a", "d", "d", "e"}
	removeDuplicates(&input)
	fmt.Println(input)
}
