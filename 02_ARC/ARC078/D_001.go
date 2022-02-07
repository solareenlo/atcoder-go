package main

import "fmt"

var (
	v  [][]int
	d1 []int
	d2 []int
)

func dfs1(x, y int) {
	for _, i := range v[x] {
		if i != y {
			d1[i] = d1[x] + 1
			dfs1(i, x)
		}
	}
}

func dfs2(x, y int) {
	for _, i := range v[x] {
		if i != y {
			d2[i] = d2[x] + 1
			dfs2(i, x)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	v = make([][]int, n+1)
	var x, y int
	for i := 1; i < n; i++ {
		fmt.Scan(&x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}

	d1 = make([]int, n+1)
	dfs1(1, 0)
	d2 = make([]int, n+1)
	dfs2(n, 0)

	res := 0
	for i := 1; i <= n; i++ {
		if d1[i] > d2[i] {
			res--
		} else {
			res++
		}
	}
	if res > 0 {
		fmt.Println("Fennec")
	} else {
		fmt.Println("Snuke")
	}
}
