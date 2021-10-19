package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	k -= n - 1 + (n-2)*(n-1)/2

	if 1 < n+k {
		fmt.Println(-1)
		return
	}
	fmt.Println(-k)

	for i := 1; i < n && k != 0; i++ {
		for j := i + 1; j < n+1 && k != 0; j++ {
			fmt.Println(i, j)
			k++
		}
	}
}
