package main

import "fmt"

func main() {
	c := make([]string, 16)
	for i := range c {
		fmt.Scan(&c[i])
	}
	for i := 15; i >= 0; i-- {
		fmt.Printf("%s ", c[i])
		if i%4 == 0 {
			fmt.Println()
		}
	}
}
