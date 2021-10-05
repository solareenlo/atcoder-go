package main

import "fmt"

func main() {
	c := make([]string, 3)
	for i := range c {
		fmt.Scan(&c[i])
	}

	fmt.Printf("%c%c%c\n", c[0][0], c[1][1], c[2][2])
}
