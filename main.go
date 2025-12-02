package main

import (
	"fmt"

	"nguyenvanhuong.vn/aoc2025/day1"
)

func main() {
	password, err := day1.Process("day1/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 1: secret entrance: %d\n", password)
}
