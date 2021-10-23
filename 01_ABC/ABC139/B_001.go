package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b == 1 {
		fmt.Println(0)
	} else if a >= b {
		fmt.Println(1)
	} else {
		res := 1
		res += (b - a) / (a - 1)
		if (b-a)%(a-1) != 0 {
			res++
		}
		fmt.Println(res)
	}
}
