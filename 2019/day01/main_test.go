package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testCases := []struct {
		mass int
		want int
	}{
		{mass: 12, want: 2},
		{mass: 14, want: 2},
		{mass: 1969, want: 654},
		{mass: 100756, want: 33583},
	}

	for _, tt := range testCases {
		got := CalculateFuel(tt.mass)
		if got != tt.want {
			t.Errorf("CalculateFuel(%#v) = %d, want %d", tt.mass, got, tt.want)
		}
	}
}

func TestCalculateFuelRecursive(t *testing.T) {
	testCases := []struct {
		mass int
		want int
	}{
		{mass: 12, want: 2},
		{mass: 14, want: 2},
		{mass: 1969, want: 966},
		{mass: 100756, want: 50346},
	}

	for _, tt := range testCases {
		got := CalculateFuelRecursive(tt.mass)
		if got != tt.want {
			t.Errorf("CalculateFuelRecursive(%#v) = %d, want %d", tt.mass, got, tt.want)
		}
	}
}
