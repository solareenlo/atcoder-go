package main

import "fmt"

func main() {
	var n, x, m int
	fmt.Scan(&n, &x, &m)

	vis := make([]int, 100001)
	for i := 0; i < m; i++ {
		vis[i] = -1
	}

	pre := make([]int, 100001)
	res, tu := 0, 0
	for n > 0 {
		n--
		vis[x] = tu
		tu++
		pre[x] = res
		res += x
		x = x * x % m
		if vis[x] >= 0 {
			be := tu - vis[x]
			res += (res - pre[x]) * (n / be)
			n %= be
		}
	}

	fmt.Println(res)
}
