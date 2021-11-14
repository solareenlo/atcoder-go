package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)

	res := 1
	for i := 1; i < 12; i++ {
		res *= l - i
		res /= i
	}
	fmt.Println(res)
}
