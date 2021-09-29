package main

import "fmt"

func main() {
	var x, t int
	fmt.Scan(&x, &t)

	res := 0
	if x > t {
		res = x - t
	}
	fmt.Println(res)
}
