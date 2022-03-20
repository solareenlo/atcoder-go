package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c == 1 {
		fmt.Println("No")
		return
	}

	v := 1
	for i := 1; i <= b; i++ {
		if a/c < v {
			fmt.Println("Yes")
			return
		}
		v *= c
	}
	fmt.Println("No")
}
