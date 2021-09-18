package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	res := x / (6 + 5) * 2
	div := x % (6 + 5)
	if div != 0 {
		if div <= 6 {
			res += 1
		} else {
			res += 2
		}
	}
	fmt.Println(res)
}
