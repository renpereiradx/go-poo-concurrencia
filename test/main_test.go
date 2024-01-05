package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	// if Sum(1, 2) != 3 {
	// 	t.Error("Sum(1, 2) should be 3")
	// }
	table := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}
	for _, v := range table {
		if v.c != Sum(v.a, v.b) {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", Sum(v.a, v.b), v.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{7, 7, 7},
	}
	for _, tt := range table {
		t.Run(fmt.Sprintf("GetMax(%d, %d)", tt.a, tt.b), func(t *testing.T) {
			if got := GetMax(tt.a, tt.b); got != tt.want {
				t.Errorf("GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	table := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
	}
	for _, tt := range table {
		t.Run(fmt.Sprintf("Fibonacci(%d)", tt.n), func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFibonacci2(t *testing.T) {
	table := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
	}
	for _, tt := range table {
		t.Run(fmt.Sprintf("Fibonacci2(%d)", tt.n), func(t *testing.T) {
			if got := Fibonacci2(tt.n); got != tt.want {
				t.Errorf("Fibonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}
