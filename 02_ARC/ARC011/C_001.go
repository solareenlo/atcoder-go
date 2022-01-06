package main

import "fmt"

func main() {
	var f, l string
	var n int
	fmt.Scan(&f, &l, &n)

	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	if f == l {
		fmt.Println(0)
		fmt.Println(f)
		fmt.Println(l)
		return
	}

	m := map[string]int{}
	for i := 0; i < n; i++ {
		m[s[i]] = i
	}
	if _, ok := m[f]; !ok {
		m[f] = n
		n++
		s = append(s, f)
	}
	if _, ok := m[l]; !ok {
		m[l] = n
		n++
		s = append(s, l)
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	dp[m[f]] = 0
	pr := make([]int, n)
	q := make([]int, 0)
	q = append(q, m[f])
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 0; i < n; i++ {
			if ^dp[i] != 0 {
				continue
			}
			k := 0
			for j := 0; j < len(s[p]); j++ {
				if s[p][j] != s[i][j] {
					k++
				}
			}
			if k == 1 {
				dp[i] = dp[p] + 1
				pr[i] = p
				q = append(q, i)
			}
		}
	}
	if dp[m[l]] < 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(dp[m[l]] - 1)
	ans := make([]string, 0)
	k := m[l]
	for k != m[f] {
		ans = append(ans, s[k])
		k = pr[k]
	}
	ans = append(ans, f)
	ans = reverseOrderString(ans)
	for i := range ans {
		fmt.Println(ans[i])
	}
}

func reverseOrderString(a []string) []string {
	n := len(a)
	res := make([]string, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
