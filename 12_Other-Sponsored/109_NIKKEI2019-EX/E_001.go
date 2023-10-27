package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		a := 0
		if i%2 == 0 {
			fmt.Print("a")
			a = 1
		}
		if i%3 == 0 {
			fmt.Print("b")
			a = 1
		}
		if i%4 == 0 {
			fmt.Print("c")
			a = 1
		}
		if i%5 == 0 {
			fmt.Print("d")
			a = 1
		}
		if i%6 == 0 {
			fmt.Print("e")
			a = 1
		}
		if a == 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
