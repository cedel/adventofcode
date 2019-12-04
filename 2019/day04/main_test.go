package main

import "testing"

func TestValidPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{input: 11111, want: false},
		{input: 111111, want: true},
		{input: 223450, want: false},
		{input: 123789, want: false},
	}

	for _, tt := range testCases {
		got := ValidPassword(tt.input)
		if got != tt.want {
			t.Errorf("ValidPassword(%#v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestValidStrictPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{input: 112233, want: true},
		{input: 123444, want: false},
		{input: 123334, want: false},
		{input: 111122, want: true},
	}

	for _, tt := range testCases {
		got := ValidStrictPassword(tt.input)
		if got != tt.want {
			t.Errorf("ValidStrictPassword(%#v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
