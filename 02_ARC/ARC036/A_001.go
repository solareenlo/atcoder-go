package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	t := make([]int, n)
	for i := range t {
		fmt.Scan(&t[i])
	}

	ok := false
	index := -1
	for i := 0; i < n-2; i++ {
		if t[i]+t[i+1]+t[i+2] < k {
			ok = true
			index = i + 3
			break
		}
	}

	if ok {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}
}
