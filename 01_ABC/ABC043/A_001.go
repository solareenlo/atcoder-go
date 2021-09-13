package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(res)
}
