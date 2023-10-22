package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	ans := make([]int, N)
	for i := 0; 5*i < N; i++ {
		fmt.Print("?")
		for j := 0; j < N; j++ {
			fmt.Print(" ")
			if j/5 == i {
				fmt.Print(pow(10, j%5))
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
		var res int
		fmt.Scan(&res)
		for j := 0; j < min(5, N-5*i); j++ {
			if res%2 > 0 {
				ans[5*i+j] = 1
			}
			if res%10 < 3 {
				res -= 10
			}
			res /= 10
		}
	}
	fmt.Print("!")
	for i := 0; i < N; i++ {
		fmt.Print(" ", ans[i])
	}
	fmt.Println()
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
