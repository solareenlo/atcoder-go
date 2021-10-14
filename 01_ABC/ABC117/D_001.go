package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	x := 0
	for bit := 40; bit >= 0; bit-- {
		cnt := 0
		for i := 0; i < n; i++ {
			cnt += (a[i] >> bit) & 1
		}
		if 2*cnt < n && (x|1<<bit) <= k {
			x |= 1 << bit
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += x ^ a[i]
	}
	fmt.Println(sum)
}
