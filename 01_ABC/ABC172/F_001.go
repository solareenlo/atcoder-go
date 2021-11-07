package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	c := 0
	for i := 3; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		c ^= x
	}

	s := a + b
	d := (s - c)
	ans := d / 2

	if d&1 != 0 || d < 0 || ans > a || ans&c != 0 {
		fmt.Println(-1)
	} else {
		for i := 1 << 61; i > 0; i >>= 1 {
			if i&c != 0 && ans+i <= a {
				ans += i
			}
		}
		if ans != 0 {
			fmt.Println(a - ans)
		} else {
			fmt.Println(-1)
		}
	}
}
