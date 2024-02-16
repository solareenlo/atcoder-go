package main

import "fmt"

// http://hydra.nat.uni-magdeburg.de/packing/csq/csq11.html
var RESULT = [][]int{{14239924, 14239924}, {42719771, 14239924}, {84329673, 15670327}, {58250656, 38112363}, {14239924, 42719771}, {85760076, 45483490}, {38112363, 58250656}, {65621783, 65621783}, {15670327, 84329673}, {45483490, 85760076}, {85760076, 85760076}}

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
