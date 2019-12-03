package utils

import "testing"

func TestToInt(t *testing.T) {
	input := "42"
	want := 42
	got := ToInt(input)
	if got != want {
		t.Errorf("ToInt(%#v) = %d, want %d", input, got, want)
	}
}
