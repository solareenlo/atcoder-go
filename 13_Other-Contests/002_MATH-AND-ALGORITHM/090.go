package main

import "fmt"

func f(x int) int {
	p := 1
	for x > 0 {
		p *= x % 10
		x /= 10
	}
	return p
}

func main() {
	var N, B int
	fmt.Scan(&N, &B)

	a := 1
	ans := 0
	for a <= N {
		b := a
		for b <= N {
			c := b
			for c <= N {
				d := c
				for d <= N {
					m := B + d
					if 1 <= m && m <= N && f(m) == d {
						ans++
					}
					d *= 7
				}
				c *= 5
			}
			b *= 3
		}
		a *= 2
	}

	m := B
	if 1 <= m && m <= N && f(m) == 0 {
		ans++
	}
	fmt.Println(ans)
}
