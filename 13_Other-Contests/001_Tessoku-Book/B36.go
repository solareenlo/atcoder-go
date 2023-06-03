package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			c++
		}
	}
	if abs(k-c)%2 != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
