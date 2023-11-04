package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	count := 0
	for A != B {
		if A > B {
			if A%B == 0 {
				count += (A / B) - 1
				A = B
			} else {
				count += (A / B)
				A %= B
			}
		} else {
			if B%A == 0 {
				count += (B / A) - 1
				B = A
			} else {
				count += (B / A)
				B %= A
			}
		}
	}
	fmt.Println(count)
}
