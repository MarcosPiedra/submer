package exercise_one

import (
	"slices"
	"strings"
	"testing"
)

func Test_RemoveDuplicates(t *testing.T) {
	// Arrange
	input := generateSlice()
	expected := []string{"a", "b", "c"}

	// Act
	removeDuplicates(&input)

	// Assert
	if !slices.Equal(input, expected) {
		t.Errorf("expected %v but got %v", expected, input)
	}
}

func Benchmark_RemoveDuplicates(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := generateSlice()
		b.StartTimer()
		removeDuplicates2(&input)
	}
}

func Benchmark_RemoveDuplicates2(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := generateSlice()
		b.StartTimer()
		removeDuplicates2(&input)
	}
}

func generateSlice() []string {

	s1 := strings.Split(strings.Repeat("a", 4), "")
	s2 := strings.Split(strings.Repeat("b", 4), "")
	s3 := strings.Split(strings.Repeat("c", 4), "")
	s4 := strings.Split(strings.Repeat("a", 4), "")

	input := append(s1, s2...)
	input = append(input, s3...)
	input = append(input, s4...)

	return input
}
