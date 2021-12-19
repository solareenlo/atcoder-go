package main

import "fmt"

func main() {
	var s, t, x int
	fmt.Scan(&s, &t, &x)

	if t < s {
		t += 24
	}
	if x < s {
		x += 24
	}
	if t > x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
