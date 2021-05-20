package main

import (
	"testing"

	"github.com/FS-Frost/go-test/math"
)

func TestSum(t *testing.T) {
	expectedSum := 55
	actualSum := math.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if actualSum == expectedSum {
		return
	}

	t.Errorf("Expected sum: %d; actual: %d", expectedSum, actualSum)
}
