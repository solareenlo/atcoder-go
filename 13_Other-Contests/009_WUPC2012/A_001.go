package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	e := []int{0, 0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335, 366}
	fmt.Println(e[c] + d - e[a] - b)
}
