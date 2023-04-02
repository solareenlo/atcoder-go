package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	ptr := 1
	var a, s, c [100009]int
	for i := 1; i <= N+1; i++ {
		if i != N+1 {
			fmt.Scan(&a[i])
		}
		for a[s[ptr-1]] > a[i] {
			ptr--
			pos := s[ptr]
			if a[s[ptr-1]] > a[i] {
				c[i-s[ptr-1]-1] += a[pos] - a[s[ptr-1]]
			} else {
				c[i-s[ptr-1]-1] += a[pos] - a[i]
			}
		}
		s[ptr] = i
		ptr++
	}
	ret := 0
	for i := N; i >= 1; i-- {
		if K < c[i] {
			ret += K * i
			K -= K
		} else {
			ret += c[i] * i
			K -= c[i]
		}
	}
	fmt.Println(ret)
}
