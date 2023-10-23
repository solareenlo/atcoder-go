package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a[0] == 'S' && b[0] == 'H' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
