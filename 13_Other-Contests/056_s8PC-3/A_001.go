package main

import "fmt"

func main() {
	var a, n, k, i, j, b int
	b = 11
	fmt.Scan(&n, &k)
	n -= 2
	if n >= b {
		a += n / b * 5
		n %= b
	}
	for i = 1; i < 6; i++ {
		for j = 1; j <= n; j++ {
			if (63*j+9*i+9)%b == k {
				a++
			}
		}
	}
	fmt.Println(a)
}
