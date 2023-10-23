package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a[0] == 'y' && a[1] == 'a' && a[2] == 'h' && a[3] == a[4] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
