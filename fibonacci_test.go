package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Fibonacci of 0", 0, 0},
		{"Fibonacci of 1", 1, 1},
		{"Fibonacci of 2", 2, 1},
		{"Fibonacci of 3", 3, 2},
		{"Fibonacci of 4", 4, 3},
		{"Fibonacci of 5", 5, 5},
		{"Fibonacci of 6", 6, 8},
		{"Fibonacci of 7", 7, 13},
		{"Fibonacci of 8", 8, 21},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Fibonacci(tc.input, false, make(map[int]int))
			if result != tc.expected {
				t.Errorf("Expected Fibonacci(%d) to be %d, but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
