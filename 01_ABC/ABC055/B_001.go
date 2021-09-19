package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 1
	for i := 1; i <= n; i++ {
		res = res * i % int(1e9+7)
	}
	fmt.Println(res)
}
