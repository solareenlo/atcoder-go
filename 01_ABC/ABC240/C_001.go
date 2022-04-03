package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	l := make([]int, 10500)
	for i := 1; i < len(l); i++ {
		l[i] = 255
	}
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		for j := 10000; j >= 0; j-- {
			if l[j] == i-1 {
				l[j+a] = i
				l[j+b] = i
			}
		}
	}

	if l[x] == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
