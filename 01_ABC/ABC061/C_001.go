package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, 100001)
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		arr[a] += b
	}

	for i := range arr {
		k -= arr[i]
		if k <= 0 {
			fmt.Println(i)
			return
		}
	}
}
