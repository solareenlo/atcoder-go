package main

import "fmt"

func main() {
	var n int
	var s, t string
	fmt.Scan(&n, &s, &t)
	s = " " + s + " "
	t = " " + t + " "

	spre := -n
	s1, s2 := 0, 0
	if s[1] != t[1] {
		s2++
	}

	N := 10000010
	tt := make([]int, N+N)
	dis := 0
	for i := 1; i <= n; i++ {
		tt[s1-s2+N]++
		dis += s1 - s2 + n
		if s[i] != s[i+1] {
			s1++
		}
		if t[i] != t[i+1] {
			s2++
		}
	}

	ans := 1 << 60
	for i := N - n; i <= n+N; i++ {
		if i%2 == 0 {
			if ans > dis {
				ans = dis
			}
		}
		spre += tt[i] << 1
		dis += spre
	}
	fmt.Println(ans)
}
