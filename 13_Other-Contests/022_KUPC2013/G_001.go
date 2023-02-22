package main

import "fmt"

func main() {
	a := 4
	fmt.Println(40)
	for i := 0; i < 40; i++ {
		for j := 0; j < 40; j++ {
			if i == j {
				fmt.Print("N")
			} else if i <= a || j <= a {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		fmt.Println()
	}
}
