package main

import "fmt"

func main() {
	var A, B, C, D float64
	fmt.Scan(&A, &B, &C, &D)
	c := '_'
	if A/B == C/D {
		c = '='
	}
	if A/B < C/D {
		c = '<'
	}
	if A/B > C/D {
		c = '>'
	}
	fmt.Println(string(c))
}
