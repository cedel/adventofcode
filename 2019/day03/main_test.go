package main

import (
	"reflect"
	"testing"
)

func TestBuildPath(t *testing.T) {
	testCases := []struct {
		instructions string
		want         WirePath
	}{
		{
			instructions: "U2,R1,D3,L1",
			want: WirePath{
				Point{X: 0, Y: 1}:  1,
				Point{X: 0, Y: 2}:  2,
				Point{X: 1, Y: 2}:  3,
				Point{X: 1, Y: 1}:  4,
				Point{X: 1, Y: 0}:  5,
				Point{X: 1, Y: -1}: 6,
				Point{X: 0, Y: -1}: 7,
			},
		},
		{
			instructions: "R3,L2,U1",
			want: WirePath{
				Point{X: 1, Y: 0}: 1,
				Point{X: 2, Y: 0}: 2,
				Point{X: 3, Y: 0}: 3,
				Point{X: 1, Y: 1}: 6,
			},
		},
	}

	for _, tt := range testCases {
		got := BuildPath(tt.instructions)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BuildPath(%#v) = %#v, want %#v", tt.instructions, got, tt.want)
		}
	}

}

func TestClosestIntersection(t *testing.T) {
	testCases := []struct {
		w1   string
		w2   string
		want int
	}{
		{
			w1:   "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			w2:   "U62,R66,U55,R34,D71,R55,D58,R83",
			want: 159,
		},
		{
			w1:   "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			w2:   "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want: 135,
		},
	}

	for _, tt := range testCases {
		w1 := BuildPath(tt.w1)
		w2 := BuildPath(tt.w2)
		got := ClosestIntersection(w1, w2)
		if got != tt.want {
			t.Errorf("ClosestIntersection(%#v, %#v) = %d, want %d", tt.w1, tt.w2, got, tt.want)
		}
	}
}

func TestQuickestIntersection(t *testing.T) {
	testCases := []struct {
		w1   string
		w2   string
		want int
	}{
		{
			w1:   "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			w2:   "U62,R66,U55,R34,D71,R55,D58,R83",
			want: 610,
		},
		{
			w1:   "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			w2:   "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want: 410,
		},
	}

	for _, tt := range testCases {
		w1 := BuildPath(tt.w1)
		w2 := BuildPath(tt.w2)
		got := QuickestIntersection(w1, w2)
		if got != tt.want {
			t.Errorf("QuickestIntersection(%#v, %#v) = %d, want %d", tt.w1, tt.w2, got, tt.want)
		}
	}
}
