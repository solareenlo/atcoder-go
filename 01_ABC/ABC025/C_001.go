package main

import "fmt"

var (
	b [2][3]int = [2][3]int{}
	c [3][2]int = [3][2]int{}
	v [3][3]int = [3][3]int{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
)

func dfs(t int) int {
	if t == 9 {
		res := 0
		for i := range b {
			for j := range b[i] {
				if v[i][j] == v[i+1][j] {
					res += b[i][j]
				}
			}
		}
		for i := range c {
			for j := range c[i] {
				if v[i][j] == v[i][j+1] {
					res += c[i][j]
				}
			}
		}
		return res
	}
	a := make([]int, 0)
	for i := range v {
		for j := range v[i] {
			if v[i][j] == -1 {
				v[i][j] = t % 2
				a = append(a, dfs(t+1))
				v[i][j] = -1
			}
		}
	}
	if t%2 == 0 {
		return maxSlice(a)
	} else {
		return minSlice(a)
	}
}

func minSlice(v []int) int {
	res := 0
	for i, e := range v {
		if i == 0 || e < res {
			res = e
		}
	}
	return res
}

func maxSlice(v []int) int {
	res := int(2e9)
	for i, e := range v {
		if i == 0 || e > res {
			res = e
		}
	}
	return res
}

func main() {
	sum := 0
	for i := range b {
		for j := range b[i] {
			fmt.Scan(&b[i][j])
			sum += b[i][j]
		}
	}
	for i := range c {
		for j := range c[i] {
			fmt.Scan(&c[i][j])
			sum += c[i][j]
		}
	}
	res := dfs(0)
	fmt.Println(res)
	fmt.Println(sum - res)
}
