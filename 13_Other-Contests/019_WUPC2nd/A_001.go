package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res := 0
	for i := 1; i <= n; i++ {
		res += i * i
	}
	fmt.Println(res % m)
}
