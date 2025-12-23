package day2

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("load cvs file", func(t *testing.T) {
		ranges, err := load("input.csv")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if r := len(ranges); r != 28 {
			t.Errorf("expected: %d, got: %d", 28, r)
		}
	})

	t.Run("count invalid IDs part 1", func(t *testing.T) {
		input := []Range{
			{Min: 11, Max: 22},
			{Min: 95, Max: 115},
			{Min: 998, Max: 1012},
		}

		total := countInvalidIDs(input)
		if total != 1142 {
			t.Errorf("expected: 1142, got: %d", total)
		}
	})

	t.Run("count invalid IDs part 2", func(t *testing.T) {
		input := []Range{
			{Min: 11, Max: 22},
			{Min: 95, Max: 115},
			{Min: 998, Max: 1012},
		}

		total := countInvalidIDs2(input)
		if total != 2252 {
			t.Errorf("expected: 2252, got: %d", total)
		}
	})
}
