package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i%2 == 0 {
				fmt.Print("#")
			} else if j == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
