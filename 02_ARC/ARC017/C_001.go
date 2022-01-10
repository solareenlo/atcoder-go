package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	k := n / 2
	m := map[int]int{}
	for i := 0; i < 1<<k; i++ {
		t := 0
		for j := 0; j < k; j++ {
			if 1&(i>>j) != 0 {
				t += a[j]
			}
		}
		m[t] += 1
	}

	res := 0
	for i := 0; i < 1<<(n-k); i++ {
		t := 0
		for j := 0; j < n-k; j++ {
			if 1&(i>>j) != 0 {
				t += a[j+k]
			}
		}
		if t <= x {
			res += m[x-t]
		}
	}
	fmt.Println(res)
}
