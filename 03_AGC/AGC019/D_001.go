package main

import "fmt"

var (
	l        = make([]int, 2002)
	r        = make([]int, 2002)
	ans  int = 1 << 60
	a, b string
	n    int
)

func solve(n int) {
	for i := 0; i < n; i++ {
		l[i] = 0
		r[i] = 0
		for b[(i-l[i]+n)%n] == 48 {
			l[i]++
		}
		for b[(i+r[i])%n] == 48 {
			r[i]++
		}
	}
	for i := 0; i < n; i++ {
		t := make([]int, 2002)
		mx := 0
		cnt := 0
		for j := 0; j < n; j++ {
			if a[j] != b[(j+i)%n] {
				t[l[j]] = max(t[l[j]], r[j])
				cnt++
			}
		}
		for j := n - 1; j >= 0; j-- {
			ans = min(ans, 2*j+mx+cnt+abs(mx-i))
			mx = max(mx, t[j])
		}
	}
}

func main() {
	fmt.Scan(&a, &b)
	n = len(a)
	ca := 0
	cb := 0
	for i := 0; i < n; i++ {
		ca += int(a[i]) ^ 48
		cb += int(b[i]) ^ 48
	}

	if cb == 0 {
		if ca == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(-1)
		}
		return
	}

	solve(n)
	a = reverseString(a)
	b = reverseString(b)
	solve(n)
	fmt.Println(ans)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
