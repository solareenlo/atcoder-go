package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)

	if p == 1 {
		var price int
		fmt.Scan(&price)
		var n int
		fmt.Scan(&n)
		fmt.Println(price * n)
	}

	if p == 2 {
		var text string
		var price int
		fmt.Scan(&text, &price)
		var n int
		fmt.Scan(&n)
		fmt.Print(text, "!\n")
		fmt.Println(price * n)
	}
}
