package main

import "fmt"

func main() {
	ans := 0
	var N int
	fmt.Scan(&N)
	cnt := make(map[int]int)
	for i := 2; i*i <= N; i++ {
		for N%i == 0 {
			N /= i
			cnt[i]++
		}
	}
	if N > 1 {
		cnt[N]++
	}

	for a, b := range cnt {
		cnt := 0
		for j := a; cnt < b; j += a {
			t := j
			for t%a == 0 {
				cnt++
				t /= a
			}
			ans = max(ans, j)
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
