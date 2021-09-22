package main

import "fmt"

func main() {
	var o, e string
	fmt.Scan(&o, &e)

	if len(o) == len(e) {
		for i := 0; i < len(o); i++ {
			fmt.Printf("%c%c", o[i], e[i])
		}
		fmt.Println()
	} else {
		c := o[len(o)-1]
		for i := 0; i < len(o)-1; i++ {
			fmt.Printf("%c%c", o[i], e[i])
		}
		fmt.Printf("%c\n", c)
	}
}
