package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		cnt := 1
		a := i
		for j := 2; j*j <= i; j++ {
			for a%j == 0 {
				a /= j
				cnt++
			}
		}
		if a > 1 {
			cnt++
		}
		fmt.Print(cnt, " ")
	}
	fmt.Println()
}
