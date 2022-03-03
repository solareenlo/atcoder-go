package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i*i <= n; i++ {
		ans += (n/i-i)/2 + 1
	}
	fmt.Println(ans % 998244353)
}
