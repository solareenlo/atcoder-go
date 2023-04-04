package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	now := 1
	l := 0
	for i := 1; i <= n; i++ {
		now *= a[i]
		for now > p {
			l++
			now /= a[l]
		}
		if now == p {
			fmt.Println("Yay!")
			return
		}
	}
	fmt.Println(":(")
}
