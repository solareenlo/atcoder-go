package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	if 2*(n-1) <= k && k%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
