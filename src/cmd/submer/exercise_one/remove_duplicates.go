package exercise_one

import "sort"

func removeDuplicates(s *[]string) {

	c := *s

	sort.SliceStable(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	prev := 1
	for curr := 1; curr < len(c); curr++ {
		if c[curr-1] != c[curr] {
			c[prev] = c[curr]
			prev++
		}
	}

	*s = c[:prev]
}

func removeDuplicates2(s *[]string) {

	c := *s

	sort.SliceStable(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	prev := 1
	for curr := 1; curr < len(c); curr++ {
		if c[curr-1] != c[curr] {
			c[prev] = c[curr]
			prev++
		}
	}

	c = append(c[:0], c[:prev]...)
}
