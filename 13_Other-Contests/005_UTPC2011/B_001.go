package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	lenA := len(a)
	sum := (lenA + 1) / 2
	for i := 0; i < (lenA+1)/2; i++ {
		if k(a[i], a[lenA-1-i]) {
			sum--
		}
	}
	fmt.Println(sum)
}

func k(q, w byte) bool {
	if q == 'i' && w == 'i' {
		return true
	}
	if q == 'w' && w == 'w' {
		return true
	}
	if q == '(' && w == ')' {
		return true
	}
	if q == ')' && w == '(' {
		return true
	}
	return false
}
