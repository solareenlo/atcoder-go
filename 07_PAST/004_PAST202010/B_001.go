package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if y == 0 {
		fmt.Println("ERROR")
	} else {
		x = x * 100 / y
		fmt.Print(x / 100)
		fmt.Print(".")
		fmt.Print(x % 100 / 10)
		fmt.Print(x % 10)
	}
}
