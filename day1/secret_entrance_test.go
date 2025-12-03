package day1

import (
	"testing"
)

func TestPassword(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rotations := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}

		expected := 3
		actual := password(rotations)

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		rotations := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
		}

		expected := 2
		actual := password(rotations)

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		rotations := []string{
			"L50",
		}

		expected := 1
		actual := password(rotations)

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 4", func(t *testing.T) {
		rotations := []string{
			"R50",
		}

		expected := 1
		actual := password(rotations)

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 5", func(t *testing.T) {
		rotations := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}

		expected := 6
		actual := password2(rotations)

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})
}
