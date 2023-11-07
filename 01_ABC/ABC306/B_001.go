package main

import "fmt"

func main() {
	var z, ans uint64
	z, ans = 1, 0
	for i := 0; i <= 63; i++ {
		var a uint64
		fmt.Scan(&a)
		ans += a * z
		z *= 2
	}
	fmt.Println(ans)
}
