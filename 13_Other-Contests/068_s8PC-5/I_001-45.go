package main

import "fmt"

func main() {
	fmt.Printf("0\n")
	for i := 0; i < 50; i++ {
		for j := 0; j < 99; j++ {
			fmt.Printf("R\n")
		}
		fmt.Printf("D\n")
		for j := 0; j < 99; j++ {
			fmt.Printf("L\n")
		}
		fmt.Printf("D\n")
	}
}
