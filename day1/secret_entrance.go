package day1

import (
	"bufio"
	"os"
	"strconv"
)

func Process() (int, error) {
	rotations, err := load("day1/input.txt")
	if err != nil {
		return -1, err
	}

	return password(rotations), nil

}

func password(rotations []string) int {
	p := 50
	count := 0

	for _, r := range rotations {
		direction, steps := parse(r)
		if direction == "L" {
			p = ((p-steps)%100 + 100) % 100
		}

		if direction == "R" {
			p = (p + steps) % 100
		}

		if p == 0 {
			count++
		}
	}

	return count
}

func parse(rotation string) (direction string, turns int) {
	direction = string(rotation[0])
	turns, _ = strconv.Atoi(rotation[1:])
	return
}

func load(path string) ([]string, error) {
	file, err := os.Open(path)
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
