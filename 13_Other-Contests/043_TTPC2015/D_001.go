package main

import (
	"fmt"
)

func isPrime(a int) bool {
	if a%2 == 0 {
		return a == 2
	}
	for i := 3; i*i <= a; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return a != 1
}

var s string
var f [128]int
var used [10]bool

func dfs(i int) int {
	if i == len(s) {
		res := 0
		for j := 0; j < len(s); j++ {
			res = res*10 + f[s[j]]
		}
		if isPrime(res) {
			return res
		} else {
			return -1
		}
	}
	if f[s[i]] != 0 {
		return dfs(i + 1)
	} else {
		for _, d := range []int{1, 3, 5, 7, 9} {
			if !used[d] {
				used[d] = true
				f[s[i]] = d
				res := dfs(i + 1)
				if res != -1 {
					return res
				}
				f[s[i]] = 0
				used[d] = false
			}
		}
		return -1
	}
}

func main() {
	fmt.Scan(&s)
	fmt.Println(dfs(0))
}
