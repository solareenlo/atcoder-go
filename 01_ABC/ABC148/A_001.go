package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	switch a + b {
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(2)
	case 5:
		fmt.Println(1)
	}
}
