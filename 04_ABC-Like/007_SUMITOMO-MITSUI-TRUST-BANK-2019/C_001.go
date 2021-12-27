package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	div := x / 100
	rem := x % 100

	ok := false
	if x < 100 {
		ok = false
	} else if rem == 0 {
		ok = true
	} else if x >= div*100 && x <= div*105 {
		ok = true
	}

	if ok {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
