package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/cedel/adventofcode/utils"
)

// Execute follows the instructions contained in the provided
// program, based on the verb (position 1) and noun (position 2)
// input values. Once the program is halted, it returns the output
// contained at position 0.
func Execute(program []int, noun int, verb int) int {
	data := make([]int, len(program))
	copy(data, program)
	data[1] = noun
	data[2] = verb

	i := 0
	for {
		switch data[i] {
		case 1:
			a := data[i+1]
			b := data[i+2]
			t := data[i+3]
			data[t] = data[a] + data[b]
			i += 4
		case 2:
			a := data[i+1]
			b := data[i+2]
			t := data[i+3]
			data[t] = data[a] * data[b]
			i += 4
		case 99:
			return data[0]
		default:
			panic("Unknown command")
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	program := []int{}
	for _, s := range strings.Split(strings.TrimSpace(string(data)), ",") {
		program = append(program, utils.ToInt(s))
	}

	output := Execute(program, 12, 2)
	fmt.Println(output)

loop:
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			output = Execute(program, n, v)
			if output == 19690720 {
				fmt.Println(n*100 + v)
				break loop
			}
		}
	}
}
