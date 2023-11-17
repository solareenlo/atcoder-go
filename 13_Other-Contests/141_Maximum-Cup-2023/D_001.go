package main

import "fmt"

func main() {
	var K int
	var S string
	fmt.Scan(&K, &S)
	X := 0
	if (K-len(S))%2 == 1 {
		X = 4
	}
	for _, c := range S {
		if c == 'A' {
			X ^= 3
		} else if c == 'C' {
			X ^= 2
		} else if c == 'G' {
			X ^= 1
		}
	}
	if X != 0 {
		fmt.Println("S")
	} else {
		fmt.Println("U")
	}
}
