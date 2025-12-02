package main

import (
	"fmt"

	"nguyenvanhuong.vn/aoc2025/day1"
)

func main() {
	password, err := day1.Process()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Secret entrance: %d\n", password)
}
