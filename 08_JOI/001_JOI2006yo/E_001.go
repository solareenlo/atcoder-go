package main

import "fmt"

var n int
var ans [21000]int

func F(p int) {
	x := 1
	for i := 0; i < n+100; i++ {
		ans[i] += x / p
		x %= p
		x *= 10
	}
}

func main() {
	var k, m, r int
	fmt.Scan(&n, &k, &m, &r)
	F(n)
	if m != 0 {
		for i := 1; i < n; i++ {
			F(i * n)
		}
	}
	for i := r + 99; i >= 0; i-- {
		ans[i] += ans[i+1] / 10
		ans[i+1] %= 10
	}
	fmt.Printf("%d.", ans[0])
	for i := 1; i <= r; i++ {
		fmt.Print(ans[i])
	}
	fmt.Println()
}
