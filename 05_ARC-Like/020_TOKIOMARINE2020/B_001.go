package main

import "fmt"

func main() {
	var a, v, b, w, t int
	fmt.Scan(&a, &v, &b, &w, &t)

	if abs(b-a) <= (v-w)*t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
