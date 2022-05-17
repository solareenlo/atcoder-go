package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		ans >>= x
		if ans&1 != 0 {
			ans++
		}
		ans++
		ans <<= x
	}
	fmt.Println(ans)
}
