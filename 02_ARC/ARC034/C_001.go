package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	b++

	m := map[int]int{}
	for ; b < a+1; b++ {
		n := b
		for d := 2; d*d <= n; d++ {
			if n%d != 0 {
				continue
			}
			t := 0
			for (n % d) == 0 {
				n /= d
				t++
			}
			m[d] += t
		}
		if n != 1 {
			m[n]++
		}
	}

	p := 1
	for _, y := range m {
		p = (p * (y + 1)) % 1_000_000_007
	}
	fmt.Println(p)
}
