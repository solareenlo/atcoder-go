package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	for i := 0; i < len(n); i++ {
		if i < 3 {
			fmt.Print(string(n[i]))
		} else {
			fmt.Print(0)
		}
	}
}
