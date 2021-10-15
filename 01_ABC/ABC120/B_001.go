package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	res := 100
	for k > 0 {
		if a%res == 0 && b%res == 0 {
			k--
		}
		res--
	}
	fmt.Println(res + 1)
}
