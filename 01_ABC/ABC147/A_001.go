package main

import "fmt"

func main() {
	a := [3]int{}
	for i := range a {
		fmt.Scan(&a[i])
	}

	if a[0]+a[1]+a[2] >= 22 {
		fmt.Println("bust")
	} else {
		fmt.Println("win")
	}
}
