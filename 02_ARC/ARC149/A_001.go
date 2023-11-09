package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	c := 0
	ansp, ansq := 0, 0
	for i := 1; i <= n; i++ {
		c = (c*10 + 1) % m
		for j := 1; j <= 9; j++ {
			if c*j%m == 0 {
				ansp = i
				ansq = j
			}
		}
	}
	if ansp == 0 {
		fmt.Println(-1)
	} else {
		for i := 1; i <= ansp; i++ {
			fmt.Print(string(ansq ^ 48))
		}
	}
	fmt.Println()
}
