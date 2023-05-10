package main

import "fmt"

func main() {
	fmt.Println(2500)
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			fmt.Println(1, i, j)
		}
	}
}
