package main

import "fmt"

// http://hydra.nat.uni-magdeburg.de/packing/csq/csq6.html
var RESULT = [][]int{{18768060, 18768060}, {81231940, 18768060}, {50000000, 39589353}, {18768060, 60410647}, {81231940, 60410647}, {50000000, 81231940}}

func main() {
	for _, line := range RESULT {
		for i := 0; i < len(line); i++ {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(line[i])
		}
		fmt.Println()
	}
}
