package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	mod := int(1e9 + 7)
	res := 0
	for i := k; i <= n+1; i++ {
		l := (i - 1) * i / 2
		r := n*i - l
		res += r - l + 1
		res %= mod
	}

	fmt.Println(res)
}
