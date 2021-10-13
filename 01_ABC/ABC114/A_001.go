package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x == 7 || x == 5 || x == 3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
