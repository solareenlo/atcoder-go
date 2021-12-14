package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	for i := 0; i < 4-len(n); i++ {
		fmt.Print("0")
	}
	fmt.Println(n)
}
