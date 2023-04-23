package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ans := 1
	for i := 1; i <= n+1; i++ {
		ans *= 5
	}
	ans -= 1
	ans /= 4
	fmt.Println(ans)
}
