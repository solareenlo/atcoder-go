package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(ack(m, n))
}

func ack(x, y int) int {
	if x == 0 {
		return y + 1
	}
	if x == 1 {
		return y + 2
	}
	if x == 2 {
		return 2*y + 3
	}
	if y == 0 {
		return ack(x-1, 1)
	}
	return 2*ack(x, y-1) + 3
}
