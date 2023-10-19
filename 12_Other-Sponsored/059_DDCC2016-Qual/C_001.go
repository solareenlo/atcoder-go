package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	m := make(map[int]int)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		m[gcd(a[i], k)]++
	}
	s := 0
	for iKey, iVal := range m {
		for jKey, jVal := range m {
			if iKey > jKey || iKey*jKey%k != 0 {
				continue
			} else if iKey < jKey {
				s += iVal * jVal
			} else {
				s += (iVal*iVal - iVal) / 2
			}
		}
	}
	fmt.Println(s)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
