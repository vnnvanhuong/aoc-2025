package day3

import "strconv"

func totalMaxjoltage(banks []string) int {
	total := 0

	for _, b := range banks {
		n := len(b)
		m := 0
		for p1 := 0; p1 < n; p1++ {
			for p2 := p1 + 1; p2 < n; p2++ {
				j, _ := strconv.Atoi(string(b[p1]) + string(b[p2]))
				m = max(m, j)
			}
		}

		total += m
	}

	return total
}
