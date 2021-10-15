package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	n, res := b-a, 0
	if a%2 == 0 {
		for i := 0; i <= n%4; i++ {
			res ^= b - i
		}
	} else {
		res = a
		for i := 1; i <= n%4; i++ {
			res ^= b - i + 1
		}
	}
	fmt.Println(res)
}
