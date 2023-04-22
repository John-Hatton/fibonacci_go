package main

import (
	"math/big"
	"testing"
)

//func TestFibonacci(t *testing.T) {
//	testCases := []struct {
//		name     string
//		input    uint64
//		expected *big.Int
//	}{
//		{"Fibonacci of 0", 0, big.NewInt(0)},
//		{"Fibonacci of 1", 1, big.NewInt(1)},
//		{"Fibonacci of 2", 2, big.NewInt(1)},
//		{"Fibonacci of 3", 3, big.NewInt(2)},
//		{"Fibonacci of 4", 4, big.NewInt(3)},
//		{"Fibonacci of 5", 5, big.NewInt(5)},
//		{"Fibonacci of 6", 6, big.NewInt(8)},
//		{"Fibonacci of 7", 7, big.NewInt(13)},
//		{"Fibonacci of 8", 8, big.NewInt(21)},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			result := Fibonacci(tc.input, false, make(map[uint64]*big.Int))
//			if result != tc.expected {
//				t.Errorf("Expected Fibonacci(%d) to be %d, but got %d", tc.input, tc.expected, result)
//			}
//		})
//	}
//}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint64
		expected *big.Int
	}{
		{"Fibonacci of 0", 0, big.NewInt(0)},
		{"Fibonacci of 1", 1, big.NewInt(1)},
		{"Fibonacci of 2", 2, big.NewInt(1)},
		{"Fibonacci of 3", 3, big.NewInt(2)},
		{"Fibonacci of 4", 4, big.NewInt(3)},
		{"Fibonacci of 5", 5, big.NewInt(5)},
		{"Fibonacci of 6", 6, big.NewInt(8)},
		{"Fibonacci of 7", 7, big.NewInt(13)},
		{"Fibonacci of 8", 8, big.NewInt(21)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Fibonacci(tc.input, false, make(map[uint64]*big.Int))
			if result.Cmp(tc.expected) != 0 {
				t.Errorf("Expected Fibonacci(%d) to be %d, but got %d", tc.input, tc.expected, result)
			}
		})
	}
}
