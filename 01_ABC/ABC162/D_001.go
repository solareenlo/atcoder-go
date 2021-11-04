package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = int(s[i]) % 3
	}

	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		cnt[a[i]]++
	}

	res := 1
	for i := 0; i < 3; i++ {
		res *= cnt[i]
	}

	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			k := j + j - i
			if k < n {
				if a[i] == a[j] {
					continue
				}
				if a[j] == a[k] {
					continue
				}
				if a[k] == a[i] {
					continue
				}
				res--
			}
		}
	}

	fmt.Println(res)
}
