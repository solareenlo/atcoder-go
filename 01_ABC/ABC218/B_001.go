package main

import "fmt"

func main() {
	p := make([]int, 26)
	for i := range p {
		fmt.Scan(&p[i])
	}

	for i := range p {
		fmt.Print(string('a' + p[i] - 1))
	}
	fmt.Println()
}
