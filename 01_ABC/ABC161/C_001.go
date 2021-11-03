package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := 1 << 60
	cnt := 0
	for {
		n = min(n%k, abs(n-k))
		if res > n {
			res = n
			cnt = 0
		} else {
			cnt++
		}
		if cnt == 2 {
			break
		}
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
