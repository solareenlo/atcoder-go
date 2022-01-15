package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	div := divList(n)
	sum := 0
	for k, _ := range div {
		sum += k
	}
	sum -= n

	if sum == n {
		fmt.Println("Perfect")
	} else if sum > n {
		fmt.Println("Abundant")
	} else {
		fmt.Println("Deficient")
	}
}

func divList(n int) map[int]struct{} {
	div := map[int]struct{}{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			div[i] = struct{}{}
			if i*i != n {
				div[n/i] = struct{}{}
			}
		}
	}
	return div
}
