package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0.0
	var x float64
	var u string
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &u)
		if u == "JPY" {
			res += x
		} else {
			res += x * 380000.0
		}
	}
	fmt.Println(res)
}
