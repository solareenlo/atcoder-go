package main

import (
	"fmt"
	"strconv"
)

var v25 []int
var um map[int]int

func main() {
	var n int
	fmt.Scan(&n)
	v25 = make([]int, 0)
	for i := 1; i <= 18; i++ {
		s := ""
		for j := 0; j < i; j++ {
			if j%2 != 0 {
				s += "2"
			} else {
				s += "5"
			}
		}
		tmp, _ := strconv.Atoi(s)
		v25 = append(v25, tmp)
		s = ""
		for j := 0; j < i; j++ {
			if j%2 != 0 {
				s += "5"
			} else {
				s += "2"
			}
		}
		tmp, _ = strconv.Atoi(s)
		v25 = append(v25, tmp)
	}
	um = make(map[int]int)
	fmt.Println(dfs(n))
}

func dfs(n int) int {
	if n == 0 {
		return 0
	}
	if _, ok := um[n]; ok {
		return um[n]
	}
	ret := n
	for i := 0; i < len(v25); i++ {
		a := n / v25[i]
		b := n % v25[i]
		ret = min(ret, b+dfs(a))
	}
	um[n] = ret
	return um[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
