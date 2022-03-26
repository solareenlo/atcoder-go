package main

import "fmt"

var (
	n  int
	a  = [51]int{}
	b  = [51]int{}
	is = [51]int{}
)

func Check() bool {
	for i := 1; i <= n; i++ {
		f := [51]bool{}
		f[a[i]] = true
		for k := 50; k > 0; k-- {
			if is[k] != 0 {
				for j := a[i]; j >= k; j-- {
					if f[j] {
						f[j%k] = true
					}
				}
			}
		}
		if !f[b[i]] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&b[i])
	}

	for i := 1; i <= 50; i++ {
		is[i] = 1
	}
	for i := 50; i > 0; i-- {
		is[i] = 0
		if !Check() {
			is[i] = 1
		}
	}

	ans := 0
	for i := 1; i <= 50; i++ {
		if is[i] != 0 {
			ans |= 1 << i
		}
	}
	if !Check() {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
