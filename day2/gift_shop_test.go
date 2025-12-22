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

	t.Run("count invalid IDs", func(t *testing.T) {
		input := []Range{
			{Min: 11, Max: 22},
			{Min: 95, Max: 115},
			{Min: 998, Max: 1012},
		}

		total := countInvalidIDs(input)
		if total != 1132 {
			t.Errorf("expected: 1132, got: %d", total)
		}
	})
}
