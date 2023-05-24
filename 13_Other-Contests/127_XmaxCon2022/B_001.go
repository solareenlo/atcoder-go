package main

import "fmt"

func main() {
	var n, k, p int
	var s string
	fmt.Scan(&n, &k, &p, &s)

	pr := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			pr[i] += p
		} else {
			pr[i] += 100 - p
		}
		if i != 0 {
			pr[i] += pr[i-1]
		}
	}
	for i := 1; i < k+1; i++ {
		c := min(2*i-1, n)
		if pr[c-1]*2 > 100*c {
			fmt.Print("+")
		} else if pr[c-1]*2 < 100*c {
			fmt.Print("-")
		} else {
			fmt.Print(0)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
