package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() {
	banks, err := load("day3/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to load input: %w", err))
	}

	fmt.Println("-> part 1: ", totalMaxjoltage(banks))
}

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

func load(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
