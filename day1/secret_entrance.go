package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type direction string

const (
	Left  direction = "L"
	Right direction = "R"
)

func Solve() {
	rotations, err := load("day1/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to load input: %w", err))
	}

	fmt.Println("-> part 1: ", password(rotations))
	fmt.Println("-> part 2: ", password2(rotations))
}

func password(rotations []string) int {
	p := 50
	count := 0

	for _, r := range rotations {
		direction, steps := parse(r)
		if direction == Left {
			p = ((p-steps)%100 + 100) % 100
		}

		if direction == Right {
			p = (p + steps) % 100
		}

		if p == 0 {
			count++
		}
	}

	return count
}

func password2(rotations []string) int {
	p := 50
	count := 0

	for _, r := range rotations {
		direction, steps := parse(r)
		if direction == Left {
			if p == 0 {
				count += steps / 100
			}

			if p != 0 && steps >= p {
				count += 1 + (steps-p)/100
			}

			p = ((p-steps)%100 + 100) % 100
		}

		if direction == Right {
			count += (p + steps) / 100
			p = (p + steps) % 100
		}
	}

	return count
}

func parse(rotation string) (d direction, turns int) {
	d = direction(string(rotation[0]))
	turns, _ = strconv.Atoi(rotation[1:])
	return
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
