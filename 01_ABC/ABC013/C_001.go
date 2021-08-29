package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h, a, b, c, d, e int
	fmt.Scan(&n, &h, &a, &b, &c, &d, &e)
	res := int(1e16)
	for i := 0; i <= n; i++ {
		j := sort.Search(n-i, func(j int) bool {
			return h+b*i+d*j-(n-i-j)*e > 0
		})
		res = min(res, a*i+c*j)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
