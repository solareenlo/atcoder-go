package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	sum := a + b + c + d
	ok := false
	if sum-a == a {
		ok = true
	}
	if sum-b == b {
		ok = true
	}
	if sum-c == c {
		ok = true
	}
	if sum-d == d {
		ok = true
	}
	if c+d == a+b {
		ok = true
	}
	if b+d == a+c {
		ok = true
	}
	if b+c == a+d {
		ok = true
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
