package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	flag := -1
	for i := 1; i <= 15; i++ {
		ans := 1
		for j := 0; j < i; j++ {
			ans *= i
		}
		if ans == n {
			flag = i
			break
		}
	}
	fmt.Println(flag)
}
