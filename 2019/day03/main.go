package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/cedel/adventofcode/utils"
)

// Point represents a point on the infinite wire board
type Point struct {
	X int
	Y int
}

// Distance calculates the Manhattan distance of a Point
// from the origin point at 0,0
func (p Point) Distance() float64 {
	return math.Abs(float64(p.X)) + math.Abs(float64(p.Y))
}

// WirePath represents a path of a wire by recording each
// Point it goes through and setting the number of steps
// that correspond to the first time the wire goes through
// that specific point.
type WirePath map[Point]int

// BuildPath creates a new WirePath by parsing a given set of
// instructions.
func BuildPath(w string) WirePath {
	m := make(map[Point]int)
	x := 0
	y := 0
	s := 0
	for _, c := range strings.Split(w, ",") {
		direction := c[0]
		steps := utils.ToInt(c[1:])
		for i := 0; i < steps; i++ {
			switch direction {
			case 'U':
				y++
			case 'D':
				y--
			case 'R':
				x++
			case 'L':
				x--
			}

			s++
			p := Point{X: x, Y: y}
			if _, ok := m[p]; !ok {
				m[p] = s
			}
		}
	}
	return m
}

// ClosestIntersection determines the closest point of intersection
// of the two wires, based on the given paths and returns the
// Manhattan distance from the origin point 0,0
func ClosestIntersection(w1 WirePath, w2 WirePath) int {
	distance := math.Inf(1)

	for p, v := range w1 {
		if v > 0 && w2[p] > 0 && p.Distance() < distance {
			distance = p.Distance()
		}
	}
	return int(distance)
}

// QuickestIntersection determines the point of intersection of the
// two wires that requires the fewer total steps and returns the
// sum of steps required by the two wires starting from the origin
// point 0,0
func QuickestIntersection(w1 WirePath, w2 WirePath) int {
	totalSteps := math.Inf(1)

	for p, v := range w1 {
		if v > 0 && w2[p] > 0 {
			steps := float64(v + w2[p])
			if steps < totalSteps {
				totalSteps = steps
			}
		}
	}
	return int(totalSteps)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	w1 := BuildPath(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	w2 := BuildPath(strings.TrimSpace(scanner.Text()))

	fmt.Println(ClosestIntersection(w1, w2))
	fmt.Println(QuickestIntersection(w1, w2))
}
