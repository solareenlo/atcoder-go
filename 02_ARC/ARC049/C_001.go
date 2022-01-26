package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	x := make([]int, 110)
	y := make([]int, 110)
	for i := 0; i < a; i++ {
		fmt.Scan(&x[i], &y[i])
		x[i]--
		y[i]--
	}

	var b int
	fmt.Scan(&b)
	p := make([]int, 15)
	q := make([]int, 15)
	for i := 0; i < b; i++ {
		fmt.Scan(&p[i], &q[i])
		p[i]--
		q[i]--
	}

	ret := 0
	vis := make([]bool, 110)
	ok := make([]bool, 110)
	for i := 0; i < 1<<b; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			vis[j] = false
		}
		for true {
			for j := 0; j < n; j++ {
				ok[j] = !vis[j]
			}
			for j := 0; j < a; j++ {
				if !vis[y[j]] {
					ok[x[j]] = false
				}
			}
			for j := 0; j < b; j++ {
				if (i&(1<<j)) != 0 && !vis[p[j]] {
					ok[q[j]] = false
				}
				if i&(1<<j) == 0 {
					ok[p[j]] = false
				}
			}
			found := false
			for j := 0; j < n; j++ {
				if ok[j] {
					vis[j] = true
					cnt++
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		ret = max(ret, cnt)
	}
	fmt.Println(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
