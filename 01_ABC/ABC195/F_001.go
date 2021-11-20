package main

import "fmt"

var Prime = [20]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	dp := [1 << 20]int{}
	dp[0] = 1
	for x := a; x <= b; x++ {
		bit := 0
		for i := 0; i < 20; i++ {
			if x%Prime[i] == 0 {
				bit |= 1 << i
			}
		}
		for i := 0; i < 1<<20; i++ {
			if i&bit == 0 {
				dp[i|bit] += dp[i]
			}
		}
	}

	res := 0
	for i := 0; i < 1<<20; i++ {
		res += dp[i]
	}

	fmt.Println(res)
}
