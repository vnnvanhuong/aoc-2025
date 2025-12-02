package day1

import (
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	Left  direction = "L"
	Right direction = "R"
)

func Process(path string) (int, error) {
	rotations, err := load(path)
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

func parse(rotation string) (d direction, turns int) {
	d = direction(string(rotation[0]))
	turns, _ = strconv.Atoi(rotation[1:])
	return
}

func load(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
