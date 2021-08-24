package main

import "fmt"

var mod int = 10007

func main() {
	var n int
	fmt.Scan(&n)

	a := [3]int{0, 0, 1}
	res := 0
	for i := 3; i < n; i++ {
		res = (a[0] + a[1] + a[2]) % mod
		a[0] = a[1]
		a[1] = a[2]
		a[2] = res
	}
	if n == 1 || n == 2 {
		res = 0
	}
	if n == 3 {
		res = 1
	}
	fmt.Println(res % mod)
}
