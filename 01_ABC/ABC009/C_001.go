package main

import "fmt"

func swap(s string, i, j int) string {
	return s[:i] + s[j:j+1] + s[i+1:j] + s[i:i+1] + s[j+1:]
}

func main() {
	var n, k int
	var s, t, u string
	fmt.Scan(&n, &k, &s)
	t = s
	for i := 0; i < n; i++ {
		x := i
		for j := i + 1; j < n; j++ {
			if t[x] > t[j] {
				u = swap(t, i, j)
				cnt := 0
				for l := 0; l < n; l++ {
					if s[l] != u[l] {
						cnt++
					}
				}
				if cnt <= k {
					x = j
				}
			}
		}
		if x > i {
			t = swap(t, i, x)
		}
	}
	fmt.Println(t)
}
