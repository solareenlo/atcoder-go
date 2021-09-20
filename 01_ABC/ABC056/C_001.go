package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	res, sum := 0, 0
	for i := 0; i <= x; i++ {
		sum += i
		if sum >= x {
			res = i
			break
		}
	}
	fmt.Println(res)
}
