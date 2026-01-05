package day3

import "testing"

func TestTotalMaxjoltage(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		banks := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}

		expected := 357
		actual := totalMaxjoltage(banks)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})
}
