package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	cnt := [1048576]int{}
	for i := 2; i < r+1; i++ {
		if cnt[i] != 0 {
			continue
		}
		for j := 0; j < r+1; j += i {
			cnt[j]++
		}
		for j := i * i; j < r+1; j += i * i {
			cnt[j] = -1000000007
		}
	}

	res := 0
	for i := 2; i < r+1; i++ {
		if cnt[i] < 0 {
			continue
		}
		cc := (r / i) - ((l - 1) / i)
		if cnt[i]%2 != 0 {
			res += (cc * (cc - 1)) / 2
		} else {
			res -= (cc * (cc - 1)) / 2
		}
	}

	for i := max(2, l); i <= r; i++ {
		res -= (r/i - 1)
	}
	fmt.Println(2 * res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
