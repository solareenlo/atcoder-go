package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	res := "Strong"
	if x[0] == x[1] && x[0] == x[2] && x[0] == x[3] {
		res = "Weak"
	}
	if f(x[0], x[1]) && f(x[1], x[2]) && f(x[2], x[3]) {
		res = "Weak"
	}
	fmt.Println(res)
}

func f(a, b byte) bool {
	if a+1 == b {
		return true
	}
	return a == '9' && b == '0'
}
