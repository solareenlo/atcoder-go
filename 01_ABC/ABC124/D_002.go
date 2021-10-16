package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)
	num := make([]int, 0)
	now, cnt := 1, 0
	for i := 0; i < n; i++ {
		if int(s[i]) == '0'+now {
			cnt++
		} else {
			num = append(num, cnt)
			now = 1 - now
			cnt = 1
		}
	}

	if cnt != 0 {
		num = append(num, cnt)
	}
	if len(num)%2 == 0 {
		num = append(num, 0)
	}
	N := len(num)

	sum := make([]int, N+1)
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + num[i]
	}

	add, res := 2*k+1, 0
	for i := 0; i < N; i += 2 {
		l, r := i, min(i+add, N)
		res = max(res, sum[r]-sum[l])
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
