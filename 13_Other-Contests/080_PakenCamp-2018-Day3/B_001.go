package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	S := 0
	ans := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		S += a
		if S <= 2018 {
			ans++
		}
	}
	fmt.Println(ans)
}
