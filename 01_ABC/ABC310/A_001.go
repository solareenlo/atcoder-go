package main

import "fmt"

func main() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)
	minn := int(1e18)
	for i := 1; i <= n; i++ {
		var fw int
		fmt.Scan(&fw)
		minn = min(fw, minn)
	}
	fmt.Println(min(p, q+minn))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
