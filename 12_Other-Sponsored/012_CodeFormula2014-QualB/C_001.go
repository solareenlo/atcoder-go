package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)
	c := 0
	var a, b [30]int
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			c++
		}
		a[A[i]-'a']++
		b[B[i]-'a']++
	}
	m := false
	for i := 0; i < 30; i++ {
		if m || a[i] >= 2 {
			m = true
		} else {
			m = false
		}
		if a[i] != b[i] {
			fmt.Println("NO")
			return
		}
	}
	if c == 0 || c == 3 || c == 4 {
		if m {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		if c == 2 || c == 5 || c == 6 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
