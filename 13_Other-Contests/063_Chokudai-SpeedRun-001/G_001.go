package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a, res string
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		res += a
		x, _ := strconv.Atoi(res)
		res = strconv.Itoa(x % 1000000007)
	}
	fmt.Println(res)
}
