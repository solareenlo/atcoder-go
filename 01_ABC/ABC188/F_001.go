package main

import "fmt"

var (
	x, y int
	m    = map[int]int{}
)

func dfs(y int) int {
	if y == 1 {
		return abs(y - x)
	}
	if _, ok := m[y]; ok {
		return m[y]
	}
	if y%2 != 0 {
		m[y] = min(dfs(y+1)+1, dfs(y-1)+1, abs(y-x))
	} else {
		m[y] = min(dfs(y/2)+1, abs(y-x))
	}
	return m[y]
}

func main() {
	fmt.Scan(&x, &y)
	fmt.Println(dfs(y))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
