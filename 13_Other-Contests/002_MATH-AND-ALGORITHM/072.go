package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	ans := 1
	for i := 1; i <= B; i++ {
		if B/i-(A+i-1)/i >= 1 {
			ans = i
		}
	}
	fmt.Println(ans)
}
