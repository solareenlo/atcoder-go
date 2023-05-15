package main

import "fmt"

func main() {
	var a [5]int
	ans := 0
	for i := 1; i <= 4; i++ {
		fmt.Scan(&a[i])
		if a[i] == 1111 {
			ans++
		}
	}
	fmt.Println(ans)
}
