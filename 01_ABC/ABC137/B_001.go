package main

import "fmt"

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	start := x - k + 1
	cnt := k*2 - 1
	for i := 0; i < cnt; i++ {
		fmt.Print(i + start)
		if i != cnt-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
