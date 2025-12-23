package day2

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func Solve() {
	ranges, err := load("day2/input.csv")
	if err != nil {
		panic(err)
	}

	count1 := countInvalidIDs(ranges)
	fmt.Println("-> part 1", count1)

	count2 := countInvalidIDs2(ranges)
	fmt.Println("-> part 1", count2)
}

func countInvalidIDs2(ranges []Range) int {
	total := 0

	for _, r := range ranges {
		for i := r.Min; i <= r.Max; i++ {
			s := strconv.Itoa(i)
			n := len(s)

			for p := 1; p <= n/2; p++ {
				if n%p != 0 {
					continue
				}

				if isRepeatingPattern(s, p) {
					total += i
					break
				}
			}
		}
	}

	return total
}

func isRepeatingPattern(s string, p int) bool {
	pattern := s[:p]

	regexPattern := "^" + regexp.QuoteMeta(pattern) + "$"
	regex := regexp.MustCompile(regexPattern)

	for i := 0; i < len(s); i += p {
		segment := s[i : i+p]
		if !regex.MatchString(segment) {
			return false
		}
	}

	return true
}

func countInvalidIDs(input []Range) int {
	total := 0
	for _, r := range input {
		for i := r.Min; i <= r.Max; i++ {
			s := strconv.Itoa(i)
			n := len(s)
			if n%2 != 0 {
				continue
			}

			half := n / 2
			l := s[:half]
			r := s[half:]
			if l != r {
				continue
			}

			total += i
		}
	}

	return total
}

func load(filePath string) ([]Range, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := []Range{}
	for _, r := range records[0] {
		p := strings.Split(r, "-")
		m1, _ := strconv.Atoi(p[0])
		m2, _ := strconv.Atoi(p[1])
		result = append(result, Range{m1, m2})
	}

	return result, nil
}
