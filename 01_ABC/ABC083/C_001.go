package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	res := 0
	for {
		x *= 2
		res++
		if x > y {
			break
		}
	}
	fmt.Println(res)
}
