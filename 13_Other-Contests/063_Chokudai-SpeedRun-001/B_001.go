package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		res += a
	}
	fmt.Println(res)
}
