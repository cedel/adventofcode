package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/cedel/adventofcode/utils"
)

// CalculateFuel determines the fuel that corresponds to the given mass
func CalculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

// CalculateFuelRecursive determines the fuel that corresponds to the given mass
// taking into account the fuel required for the fuel itself recursively
func CalculateFuelRecursive(mass int) int {
	fuel := CalculateFuel(mass)
	if fuel < 0 {
		return 0
	}
	return fuel + CalculateFuelRecursive(fuel)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	t1 := 0
	t2 := 0
	for scanner.Scan() {
		num := utils.ToInt(scanner.Text())
		t1 += CalculateFuel(num)
		t2 += CalculateFuelRecursive(num)
	}

	fmt.Println(t1, t2)
}
